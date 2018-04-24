package generator

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/parser"
	"go/build"
	"path"
	"path/filepath"
	"strings"
)

type generatedFile struct {
	PkgName    string
	GoPkg      string
	GoProtoPkg string
	Dir        string

	FilePath string
	File     *parser.File

	TracerEnabled    bool
	GQLEnumsPrefix   string
	GQLMessagePrefix string
	Services         map[string]ServiceConfig
	Enums            map[string]EnumConfig
	Messages         map[string]MessageConfig
	Generator        *protoGenerator
}
type generator struct {
	config          *GenerateConfig
	importingParams map[*parser.File]*generatedFile
}

func (g *generator) importImportingParams(imp *parser.File) (dir, pkg string, rerr error) {
	path := filepath.Dir(imp.FilePath)
	if g.config.VendorPath != "" && strings.HasPrefix(path, g.config.VendorPath) {
		relPath, err := filepath.Rel(g.config.VendorPath, path)
		if err != nil {
			return "", "", errors.Wrap(err, "failed to resolve import relative path")
		}
		outDir, err := filepath.Abs(filepath.Join(g.config.Imports.OutputPath, relPath))
		if err != nil {
			return "", "", errors.Wrap(err, "failed to resolve import absolute path")
		}
		pkg, err = resolveGoPkg(g.config.VendorPath, outDir)
		if err != nil {
			return "", "", errors.Wrap(err, "failed to resolve import Go package")
		}
		return outDir, pkg, nil
	}
	if !strings.HasPrefix(path, filepath.Join(build.Default.GOPATH, "src")) {
		return "", "", errors.New("import File is outside GOPATH directory:")
	}
	relPath, err := filepath.Rel(filepath.Join(build.Default.GOPATH, "src"), path)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve import relative path")
	}
	outDir, err := filepath.Abs(filepath.Join(g.config.Imports.OutputPath, relPath))
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve import absolute path")
	}
	pkg, err = resolveGoPkg(g.config.VendorPath, outDir)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve import Go package")
	}
	return outDir, pkg, nil
}

func (g *generator) generate() error {
	var parsedFiles = new([]*parser.File)
	// Resolving what to generate
	var filesToGenerate []*generatedFile
	for _, cfg := range g.config.Protos {
		paths := mergePathsConfig(g.config.Paths, cfg.Paths)
		aliases := mergeAlieses(g.config.Imports.Aliases, cfg.ImportsAliases)
		file, err := parser.ParseFile(parsedFiles, aliases, paths, cfg.ProtoPath, true)
		if err != nil {
			return errors.Wrap(err, "failed to parse File")
		}
		filename := path.Base(file.FilePath)
		if strings.HasSuffix(filename, ".proto") {
			filename = strings.TrimSuffix(filename, ".proto")
		}
		outDir, err := filepath.Abs(mergeStringsConfig(g.config.OutputPath, cfg.OutputPath))
		if err != nil {
			return errors.Wrap(err, "failed to resolve File absolute path")
		}
		pkg, err := resolveGoPkg(g.config.VendorPath, outDir)
		if err != nil {
			return errors.Wrap(err, "failed to resolve go pkg")
		}
		goProtoPkg, err := resolveGoPkg(g.config.VendorPath, filepath.Dir(file.FilePath))
		if err != nil {
			return errors.Wrap(err, "failed to resolve go pkg")
		}
		filesToGenerate = append(filesToGenerate, &generatedFile{
			GoPkg:            pkg,
			PkgName:          mergeStringsConfig(g.config.OutputPkg, cfg.OutputPkg),
			Dir:              outDir,
			GoProtoPkg:       goProtoPkg,
			FilePath:         path.Join(outDir, filename) + ".go",
			File:             file,
			GQLEnumsPrefix:   cfg.GQLEnumsPrefix,
			GQLMessagePrefix: cfg.GQLMessagePrefix,
			Services:         cfg.Services,
			Enums:            cfg.Enums,
			TracerEnabled:    g.config.Tracer,
			Messages:         cfg.Messages,
		})
		// Resolving same package files
		for _, imp := range file.Imports {
			filename := path.Base(imp.FilePath)
			if strings.HasSuffix(filename, ".proto") {
				filename = strings.TrimSuffix(filename, ".proto")
			}
			var goProtoPkg string
			if v, ok := g.config.Imports.Settings[imp.FilePath]; ok && v.GoPackage != "" {
				goProtoPkg = v.GoPackage
			} else {
				v, err := resolveGoPkg(g.config.VendorPath, filepath.Dir(imp.FilePath))
				if err != nil {
					return errors.Wrap(err, "failed to resolve go pkg")
				}
				goProtoPkg = v
			}
			s := g.config.Imports.Settings[imp.FilePath]
			if isSamePackage(imp, file) {
				filesToGenerate = append(filesToGenerate, &generatedFile{
					GoPkg:            pkg,
					PkgName:          mergeStringsConfig(g.config.OutputPkg, cfg.OutputPkg),
					GoProtoPkg:       goProtoPkg,
					Dir:              outDir,
					FilePath:         path.Join(outDir, filename) + ".go",
					File:             imp,
					GQLEnumsPrefix:   s.GQLEnumsPrefix,
					GQLMessagePrefix: s.GQLMessagePrefix,
					Services:         s.Services,
					Enums:            s.Enums,
					TracerEnabled:    g.config.Tracer,
				})
				continue
			}
			dir, pkg, err := g.importImportingParams(imp)
			if err != nil {
				return errors.Wrap(err, "failed to resolve import importing params")
			}
			filesToGenerate = append(filesToGenerate, &generatedFile{
				GoPkg:            pkg,
				PkgName:          filepath.Base(pkg),
				GoProtoPkg:       goProtoPkg,
				Dir:              dir,
				FilePath:         path.Join(dir, filename) + ".go",
				File:             imp,
				GQLEnumsPrefix:   s.GQLEnumsPrefix,
				GQLMessagePrefix: s.GQLMessagePrefix,
				Services:         s.Services,
				Enums:            s.Enums,
				TracerEnabled:    g.config.Tracer,
			})
		}

	}
	for _, g := range filesToGenerate {
		g.Generator = newProtoGenerator(g, filesToGenerate)
	}
	for _, g := range filesToGenerate {
		err := g.Generator.generate()
		if err != nil {
			panic(err)
		}
	}
	var a = 1
	a++
	return nil
}
func (g *generator) normalizePaths() error {
	if g.config.VendorPath != "" {
		vp, err := filepath.Abs(g.config.VendorPath)
		if err != nil {
			return errors.Wrap(err, "failed to normalize vendor path")
		}
		g.config.VendorPath = vp
	}
	var importsSettings = make(map[string]ImportConfig)
	for key, s := range g.config.Imports.Settings {
		p, err := filepath.Abs(key)
		if err != nil {
			return errors.Wrap(err, "failed to normalize import path")
		}
		importsSettings[p] = s
	}
	g.config.Imports.Settings = importsSettings
	return nil
}
func Generate(gc *GenerateConfig) error {
	var g = generator{
		config: gc,
	}
	err := g.normalizePaths()
	if err != nil {
		return errors.Wrap(err, "failed to normalize config paths")
	}
	if err := g.generate(); err != nil {
		return errors.Wrap(err, "failed to generate schema")
	}
	return nil
}
