package proto

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

var scalarsResolvers = map[string]common.TypeResolver{
	"double": common.GqlFloat64TypeResolver,
	"float":  common.GqlFloat32TypeResolver,
	"bool":   common.GqlBoolTypeResolver,
	"string": common.GqlStringTypeResolver,

	"int64":    common.GqlInt64TypeResolver,
	"sfixed64": common.GqlInt64TypeResolver,
	"sint64":   common.GqlInt64TypeResolver,

	"int32":    common.GqlInt32TypeResolver,
	"sfixed32": common.GqlInt32TypeResolver,
	"sint32":   common.GqlInt32TypeResolver,

	"uint32":  common.GqlUInt32TypeResolver,
	"fixed32": common.GqlUInt32TypeResolver,

	"uint64":  common.GqlUInt64TypeResolver,
	"fixed64": common.GqlUInt64TypeResolver,
}

type parsedFile struct {
	File   *parser.File
	Config *ProtoFileConfig
}
type Generator struct {
	VendorPath      string
	GenerateTracers bool
	parser          parser.Parser
	parsedFiles     []parsedFile
}

func (g *Generator) prepareFile(file parsedFile) (*common.File, error) {
	pkgName, pkg, err := g.fileOutputPackage(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file output package")
	}
	enums, err := g.prepareFileEnums(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file enums")
	}
	inputs, err := g.fileInputObjects(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file input objects")
	}
	mapInputs, err := g.fileMapInputObjects(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file map input objects")
	}
	mapOutputs, err := g.fileMapOutputObjects(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file map output objects")
	}
	mapResolvers, err := g.fileInputMapResolvers(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file map resolvers")
	}
	outputMessages, err := g.fileOutputMessages(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file output messages")
	}
	messagesResolvers, err := g.fileInputMessagesResolvers(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file messages resolvers")
	}
	services, err := g.fileServices(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file services")
	}
	res := &common.File{
		PackageName:             pkgName,
		Package:                 pkg,
		Enums:                   enums,
		InputObjects:            inputs,
		InputObjectResolvers:    messagesResolvers,
		OutputObjects:           outputMessages,
		MapInputObjects:         mapInputs,
		MapInputObjectResolvers: mapResolvers,
		MapOutputObjects:        mapOutputs,
		Services:                services,
	}
	return res, nil
}

func (g *Generator) fileConfig(file *parser.File) *ProtoFileConfig {
	for _, f := range g.parsedFiles {
 		if f.File == file {
			return f.Config
		}
	}
	return nil
}

// fileGRPCSourcesPackage returns golang package of protobuf golang sources
func (g *Generator) fileGRPCSourcesPackage(file *parser.File) string {
	if file.FilePath == "/Users/yaroslavmytsyo/go/src/gitlab.egt-ua.com/back-office/backend/vendor/github.com/gogo/protobuf/protobuf/google/protobuf/timestamp.proto" {
		spew.Dump(file.GoPackage)
		spew.Dump(file.FilePath)
	}
	cfg := g.fileConfig(file)
	cfgGoPkg := cfg.GetGoPackage()

	if cfgGoPkg != "" {
		return cfg.GetGoPackage()
	}

	if file.GoPackage != "" {
		return file.GoPackage
	}
	pkg, err := resolveGoPackage(filepath.Dir(file.FilePath), g.VendorPath)
	if err != nil {
		panic(err)
	}
	return pkg
}

func (g *Generator) fileOutputPath(file *parser.File) (string, error) {
	cfg := g.fileConfig(file)
	if cfg == nil {
		absFilePath, err := filepath.Abs(file.FilePath)
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve file absolute path")
		}
		fileName := filepath.Base(file.FilePath)
		pkg, err := resolveGoPackage(filepath.Dir(absFilePath), g.VendorPath)
		var res string
		if err != nil {
			res, err = filepath.Abs(filepath.Join("./out/", "."+filepath.Dir(absFilePath), strings.TrimSuffix(fileName, ".proto")+".go"))
		} else {
			res, err = filepath.Abs(filepath.Join("./out/", "."+pkg, strings.TrimSuffix(fileName, ".proto")+".go"))
		}
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve absolute output path")
		}
		return res, nil
	}
	return filepath.Join(cfg.OutputPath, strings.TrimSuffix(filepath.Base(cfg.ProtoPath), ".proto")+".go"), nil
}

func (g *Generator) fileOutputPackage(file *parser.File) (name, pkg string, err error) {
	outPath, err := g.fileOutputPath(file)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file output path")
	}
	pkg, err = resolveGoPackage(filepath.Dir(outPath), g.VendorPath)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file go package")
	}
	return strings.Replace(filepath.Base(pkg), "-", "_", -1), pkg, nil
}

func (g *Generator) Generate() error {
	for _, pf := range g.parser.ParsedFiles() {
		cfg := g.fileConfig(pf)
		file, err := g.prepareFile(parsedFile{
			File:   pf,
			Config: cfg,
		})
		if err != nil {
			return errors.Wrap(err, "failed to prepare file for generation")
		}
		filePath, err := g.fileOutputPath(pf)
		if err != nil {
			return errors.Wrap(err, "failed to resolve file output path")
		}
		err = os.MkdirAll(filepath.Dir(filePath), 0777)
		if err != nil {
			return errors.Wrap(err, "failed to create directories for generated file")
		}
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		if err != nil {
			return errors.Wrap(err, "failed to open file")
		}
		err = common.Generate(file, f)
		if err != nil {
			return errors.Wrap(err, "failed to generate file")
		}
	}
	return nil
}

func (g *Generator) AddSourceByConfig(config *ProtoFileConfig) error {
	file, err := g.parser.Parse(config.ProtoPath, config.ImportsAliases, config.Paths)
	if err != nil {
		return errors.Wrap(err, "failed to parse proto file")
	}
	g.parsedFiles = append(g.parsedFiles, parsedFile{
		File:   file,
		Config: config,
	})
	return nil
}
