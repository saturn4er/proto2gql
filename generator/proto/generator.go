package proto

import (
	"os"
	"path/filepath"
	"strings"

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
	File           *parser.File
	Config         *ProtoFileConfig
	OutputPath     string
	OutputPkg      string
	OutputPkgName  string
	GRPCSourcesPkg string
}
type Generator struct {
	VendorPath      string
	GenerateTracers bool
	parser          parser.Parser
	ParsedFiles     []*parsedFile
}

func (g *Generator) parsedFile(file *parser.File) (*parsedFile, error) {
	for _, f := range g.ParsedFiles {
		if f.File == file {
			return f, nil
		}
	}
	outPath, err := g.fileOutputPath(nil, file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve file '%s' output path", file.FilePath)
	}
	outPkgName, outPkg, err := g.fileOutputPackage(nil, file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve file '%s' output Go package", file.FilePath)
	}
	grpcPkg, err := g.fileGRPCSourcesPackage(nil, file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve file '%s' GRPC sources Go package", file.FilePath)
	}
	res := &parsedFile{
		File:           file,
		Config:         nil,
		OutputPath:     outPath,
		OutputPkg:      outPkg,
		OutputPkgName:  outPkgName,
		GRPCSourcesPkg: grpcPkg,
	}
	g.ParsedFiles = append(g.ParsedFiles, res)
	return res, nil

}
func (g *Generator) prepareFile(file *parsedFile) (*common.File, error) {
	enums, err := g.prepareFileEnums(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file enums")
	}
	inputs, err := g.fileInputObjects(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file input objects")
	}
	mapInputs, err := g.fileMapInputObjects(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file map input objects")
	}
	mapOutputs, err := g.fileMapOutputObjects(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file map output objects")
	}
	mapResolvers, err := g.fileInputMapResolvers(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file map resolvers")
	}
	outputMessages, err := g.fileOutputMessages(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file output messages")
	}
	messagesResolvers, err := g.fileInputMessagesResolvers(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file messages resolvers")
	}
	services, err := g.fileServices(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file services")
	}
	res := &common.File{
		PackageName:             file.OutputPkgName,
		Package:                 file.OutputPkg,
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
	for _, f := range g.ParsedFiles {
		if f.File == file {
			return f.Config
		}
	}
	return nil
}

// fileGRPCSourcesPackage returns golang package of protobuf golang sources
func (g *Generator) fileGRPCSourcesPackage(cfg *ProtoFileConfig, file *parser.File) (string, error) {
	cfgGoPkg := cfg.GetGoPackage()

	if cfgGoPkg != "" {
		return cfg.GetGoPackage(), nil
	}

	if file.GoPackage != "" {
		return file.GoPackage, nil
	}
	fileDir := filepath.Dir(file.FilePath)
	pkg, err := GoPackageByPath(fileDir, g.VendorPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to resolve resolve go package of '%s'", fileDir)
	}
	return pkg, nil
}

func (g *Generator) fileOutputPath(cfg *ProtoFileConfig, file *parser.File) (string, error) {
	if cfg.GetOutputPath() == "" {
		absFilePath, err := filepath.Abs(file.FilePath)
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve file absolute path")
		}
		fileName := filepath.Base(file.FilePath)
		pkg, err := GoPackageByPath(filepath.Dir(absFilePath), g.VendorPath)
		var res string
		if err != nil {
			res, err = filepath.Abs(filepath.Join("./out/", "./"+filepath.Dir(absFilePath), strings.TrimSuffix(fileName, ".proto")+".go"))
		} else {
			res, err = filepath.Abs(filepath.Join("./out/", "./"+pkg, strings.TrimSuffix(fileName, ".proto")+".go"))
		}
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve absolute output path")
		}
		return res, nil
	}
	return filepath.Join(cfg.OutputPath, strings.TrimSuffix(filepath.Base(file.FilePath), ".proto")+".go"), nil
}

func (g *Generator) fileOutputPackage(cfg *ProtoFileConfig, file *parser.File) (name, pkg string, err error) {
	outPath, err := g.fileOutputPath(cfg, file)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file output path")
	}
	pkg, err = GoPackageByPath(filepath.Dir(outPath), g.VendorPath)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file go package")
	}
	return strings.Replace(filepath.Base(pkg), "-", "_", -1), pkg, nil
}
func (g *Generator) Generate() error {
	for _, f := range g.parser.ParsedFiles() {
		pf, err := g.parsedFile(f)
		if err != nil {
			return errors.Wrapf(err, "failed to resovle parsed file of '%s'", f.FilePath)
		}

		commonFile, err := g.prepareFile(pf)
		if err != nil {
			return errors.Wrap(err, "failed to prepare file for generation")
		}

		err = os.MkdirAll(filepath.Dir(pf.OutputPath), 0777)
		if err != nil {
			return errors.Wrap(err, "failed to create directories for generated file")
		}
		f, err := os.OpenFile(pf.OutputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		if err != nil {
			return errors.Wrap(err, "failed to open file")
		}
		err = common.Generate(commonFile, f)
		if err != nil {
			f.Close()
			return errors.Wrap(err, "failed to generate file")
		}
		err = f.Close()
		if err != nil {
			return errors.Wrap(err, "failed to close file")
		}
	}
	return nil
}

func (g *Generator) AddSourceByConfig(config *ProtoFileConfig) error {
	file, err := g.parser.Parse(config.ProtoPath, config.ImportsAliases, config.Paths)
	if err != nil {
		return errors.Wrap(err, "failed to parse proto file")
	}
	outPath, err := g.fileOutputPath(config, file)
	if err != nil {
		return errors.Wrapf(err, "failed to resolve file '%s' output path", file.FilePath)
	}
	outPkgName, outPkg, err := g.fileOutputPackage(config, file)
	if err != nil {
		return errors.Wrapf(err, "failed to resolve file '%s' output Go package", file.FilePath)
	}
	grpcPkg, err := g.fileGRPCSourcesPackage(config, file)
	if err != nil {
		return errors.Wrapf(err, "failed to resolve file '%s' GRPC sources Go package", file.FilePath)
	}
	g.ParsedFiles = append(g.ParsedFiles, &parsedFile{
		File:           file,
		Config:         config,
		OutputPath:     outPath,
		OutputPkg:      outPkg,
		OutputPkgName:  outPkgName,
		GRPCSourcesPkg: grpcPkg,
	})
	return nil
}
