package generator

import (
	"go/build"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/parser"
)

type gqlProtoDerivativeFile struct {
	GoProtoPkg string

	OutGoPkgName string
	OutGoPkg     string

	OutDir      string
	OutFilePath string

	ProtoFile *parser.File

	TracerEnabled    bool
	GQLEnumsPrefix   string
	GQLMessagePrefix string
	Services         map[string]ServiceConfig
	Messages         map[string]MessageConfig
	Generator        *gqlProtoDerivativeFileGenerator
}
type generator struct {
	config *GenerateConfig
}

func (g *generator) importDirAndPkg(filePath string) (dir, pkg string, rerr error) {
	pth := filepath.Dir(filePath)
	var prefixPath string
	if g.config.VendorPath != "" && strings.HasPrefix(pth, g.config.VendorPath) {
		prefixPath = g.config.VendorPath
	} else if strings.HasPrefix(pth, filepath.Join(build.Default.GOPATH, "src")) {
		prefixPath = filepath.Join(build.Default.GOPATH, "src")
	} else {
		return "", "", errors.New("import File is outside GOPATH directory")
	}
	relPath, err := filepath.Rel(prefixPath, pth)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve import relative path")
	}
	outDir, err := filepath.Abs(filepath.Join(g.config.Imports.OutputPath, relPath))
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve import absolute path")
	}
	pkg, err = g.resolveGoPkg(outDir)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve import Go package")
	}
	return outDir, pkg, nil
}

func (g *generator) generateSchemas(protos map[*ProtoConfig]*gqlProtoDerivativeFile) error {
	for i, sc := range g.config.Schemas {
		err := generateSchema(g.config, sc, protos)
		if err != nil {
			return errors.Wrapf(err, "failed to generated %d schema", i)
		}
	}
	return nil
}

func (g *generator) generate() error {
	p := parser.New(g.config.Imports.Aliases, g.config.Paths)
	// Resolving what to generate
	var filesToGenerate []*gqlProtoDerivativeFile
	var generatedProtos = make(map[*ProtoConfig]*gqlProtoDerivativeFile)

	for _, cfg := range g.config.Protos {
		p.Paths = mergePathsConfig(g.config.Paths, cfg.Paths)
		p.ImportAliases = mergeAliases(g.config.Imports.Aliases, cfg.ImportsAliases)
		file, err := p.Parse(cfg.ProtoPath)
		if err != nil {
			return errors.Wrap(err, "failed to parse File")
		}
		filename := strings.TrimSuffix(path.Base(file.FilePath), ".proto") + ".go"
		outDir := mergeStringsConfig(g.config.OutputPath, cfg.OutputPath)
		pkg, err := g.resolveGoPkg(outDir)
		if err != nil {
			return errors.Wrap(err, "failed to resolve go pkg")
		}
		goProtoPkg, err := g.resolveGoPkg(filepath.Dir(file.FilePath))
		if err != nil {
			return errors.Wrap(err, "failed to resolve go pkg")
		}
		f := &gqlProtoDerivativeFile{
			OutGoPkg:         pkg,
			OutGoPkgName:     mergeStringsConfig(g.config.OutputPkg, cfg.OutputPkg),
			OutDir:           outDir,
			GoProtoPkg:       goProtoPkg,
			OutFilePath:      path.Join(outDir, filename),
			ProtoFile:        file,
			GQLEnumsPrefix:   cfg.GQLEnumsPrefix,
			GQLMessagePrefix: cfg.GQLMessagePrefix,
			Services:         cfg.Services,
			TracerEnabled:    g.config.Tracer,
			Messages:         cfg.Messages,
		}
		filesToGenerate = append(filesToGenerate, f)
		generatedProtos[cfg] = f

		for _, imp := range file.Imports {
			filename := strings.TrimSuffix(path.Base(imp.FilePath), ".proto") + ".go"
			var goProtoPkg string
			if v, ok := g.config.Imports.Settings[imp.FilePath]; ok && v.GoPackage != "" {
				goProtoPkg = v.GoPackage
			} else {
				v, err := g.resolveGoPkg(filepath.Dir(imp.FilePath))
				if err != nil {
					return errors.Wrap(err, "failed to resolve go pkg")
				}
				goProtoPkg = v
			}
			s := g.config.Imports.Settings[imp.FilePath]
			if isSamePackage(imp, file) {
				filesToGenerate = append(filesToGenerate, &gqlProtoDerivativeFile{
					OutGoPkg:         pkg,
					OutGoPkgName:     mergeStringsConfig(g.config.OutputPkg, cfg.OutputPkg),
					GoProtoPkg:       goProtoPkg,
					OutDir:           outDir,
					OutFilePath:      path.Join(outDir, filename),
					ProtoFile:        imp,
					GQLEnumsPrefix:   s.GQLEnumsPrefix,
					GQLMessagePrefix: s.GQLMessagePrefix,
					Services:         nil,
					TracerEnabled:    g.config.Tracer,
				})
				continue
			}
			dir, pkg, err := g.importDirAndPkg(imp.FilePath)
			if err != nil {
				return errors.Wrap(err, "failed to resolve import importing params")
			}
			filesToGenerate = append(filesToGenerate, &gqlProtoDerivativeFile{
				OutGoPkg:         pkg,
				OutGoPkgName:     filepath.Base(pkg),
				GoProtoPkg:       goProtoPkg,
				OutDir:           dir,
				OutFilePath:      path.Join(dir, filename),
				ProtoFile:        imp,
				GQLEnumsPrefix:   s.GQLEnumsPrefix,
				GQLMessagePrefix: s.GQLMessagePrefix,
				Services:         nil,
				TracerEnabled:    g.config.Tracer,
			})
		}

	}
	for _, gf := range filesToGenerate {
		gf.Generator = newProtoGenerator(gf, filesToGenerate)
	}
	for _, g := range filesToGenerate {
		err := g.Generator.generate()
		if err != nil {
			return errors.Wrap(err, "failed to generate file")
		}
	}
	return g.generateSchemas(generatedProtos)
}

func (g *generator) resolveGoPkg(dir string) (string, error) {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	if g.config.VendorPath != "" && strings.HasPrefix(absPath, g.config.VendorPath) {
		pkg, err := filepath.Rel(g.config.VendorPath, absPath)
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve dir vendor relative path")
		}
		return pkg, nil
	} else if !strings.HasPrefix(absPath, filepath.Join(build.Default.GOPATH, "src")) {
		return "", errors.New("dir is outside GOPATH")
	}
	pkg, err := filepath.Rel(filepath.Join(build.Default.GOPATH, "src"), absPath)
	if err != nil {
		return "", errors.Wrap(err, "failed to resolve dir relative path")
	}
	return pkg, nil
}

func (g *generator) normalizeConfigPaths() error {
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
	if g.config.OutputPath != "" {
		vp, err := filepath.Abs(g.config.OutputPath)
		if err != nil {
			return errors.Wrap(err, "failed to normalize output path")
		}
		g.config.OutputPath = vp
	}
	for _, cfg := range g.config.Protos {
		if cfg.OutputPath != "" {
			vp, err := filepath.Abs(cfg.OutputPath)
			if err != nil {
				return errors.Wrap(err, "failed to normalize proto output path")
			}
			cfg.OutputPath = vp
		}
	}
	return nil
}

func Generate(gc *GenerateConfig) error {
	var g = generator{
		config: gc,
	}
	err := g.normalizeConfigPaths()
	if err != nil {
		return errors.Wrap(err, "failed to normalize config paths")
	}
	if err := g.generate(); err != nil {
		return errors.Wrap(err, "failed to generate schema")
	}
	return nil
}
