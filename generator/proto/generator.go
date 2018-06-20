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
	File   *parser.File
	Config *ProtoConfig
}
type Generator struct {
	vendorPath  string
	parser      parser.Parser
	parsedFiles []parsedFile
}

func (g *Generator) setVendorPath(path string) {
	g.vendorPath = path
}

func (g *Generator) prepareFileEnums(file parsedFile) ([]common.Enum, error) {
	var res []common.Enum
	for _, enum := range file.File.Enums {
		vals := make([]common.EnumValue, len(enum.Values))
		for i, value := range enum.Values {
			vals[i] = common.EnumValue{
				Name:    value.Name,
				Value:   value.Value,
				Comment: value.QuotedComment,
			}
		}
		res = append(res, common.Enum{
			VariableName: g.enumVariable(enum),
			GraphQLName:  g.enumGraphQLName(enum),
			Comment:      enum.QuotedComment,
			Values:       vals,
		})
	}
	return res, nil
}

func (g *Generator) enumGraphQLName(res *parser.Enum) string {
	cfg := g.fileConfig(res.File)
	return cfg.GetGQLEnumsPrefix() + res.Name
}

func (g *Generator) enumVariable(res *parser.Enum) string {
	return res.Name
}

func (g *Generator) enumTypeResolver(currentFile *parser.File, enum *parser.Enum) (common.TypeResolver, error) {
	_, pkg, err := g.fileOutputPackage(enum.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file output package")
	}
	return func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(pkg) + g.enumVariable(enum)
	}, nil
}
func (g *Generator) prepareFile(file parsedFile) (*common.File, error) {
	pkgName, pkg, err := g.fileOutputPackage(file.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file output package")
	}
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
	mapResolvers, err := g.fileMapInputObjectsResolvers(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare file map resolvers")
	}
	res := &common.File{
		PackageName:             pkgName,
		Package:                 pkg,
		Enums:                   enums,
		InputObjects:            inputs,
		MapInputObjects:         mapInputs,
		MapInputObjectResolvers: mapResolvers,
		MapOutputObjects:        mapOutputs,
	}
	return res, nil
}

func (g *Generator) fileConfig(file *parser.File) *ProtoConfig {
	for _, f := range g.parsedFiles {
		if f.File == file {
			return f.Config
		}
	}

	return nil
}

func (g *Generator) fileGRPCSourcesPackage(file *parser.File) string {
	cfg := g.fileConfig(file)
	if cfg.GetGoPackage() != "" {
		return cfg.GetGoPackage()
	}

	return file.GoPackage
}

func (g *Generator) fileOutputPath(file *parser.File) (string, error) {
	cfg := g.fileConfig(file)
	if cfg == nil {
		absFilePath, err := filepath.Abs(file.FilePath)
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve file absolute path")
		}
		fileName := filepath.Base(file.FilePath)
		res, err := filepath.Abs(filepath.Join("./out/", "."+filepath.Dir(absFilePath), strings.TrimSuffix(fileName, ".proto")+".go"))
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve absolute output path")
		}
		return res, nil
	}
	return strings.TrimSuffix(filepath.Join(cfg.OutputPath, cfg.ProtoPath), ".proto") + ".go", nil
}

func (g *Generator) fileOutputPackage(file *parser.File) (name, pkg string, err error) {
	outPath, err := g.fileOutputPath(file)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file output path")
	}
	pkg, err = resolveGoPackage(filepath.Dir(outPath), g.vendorPath)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file go package")
	}
	return filepath.Base(pkg), pkg, nil
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

func (g *Generator) AddSourceByConfig(config ProtoConfig) error {
	file, err := g.parser.Parse(config.ProtoPath, config.ImportsAliases, config.Paths)
	if err != nil {
		return errors.Wrap(err, "failed to parse proto file")
	}
	g.parsedFiles = append(g.parsedFiles, parsedFile{
		File:   file,
		Config: &config,
	})
	return nil
}
