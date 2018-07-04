package swagger2gql

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

var scalarsResolvers = map[string]graphql.TypeResolver{
	"double": graphql.GqlFloat64TypeResolver,
	"float":  graphql.GqlFloat32TypeResolver,
	"bool":   graphql.GqlBoolTypeResolver,
	"string": graphql.GqlStringTypeResolver,

	"int64":    graphql.GqlInt64TypeResolver,
	"sfixed64": graphql.GqlInt64TypeResolver,
	"sint64":   graphql.GqlInt64TypeResolver,

	"int32":    graphql.GqlInt32TypeResolver,
	"sfixed32": graphql.GqlInt32TypeResolver,
	"sint32":   graphql.GqlInt32TypeResolver,

	"uint32":  graphql.GqlUInt32TypeResolver,
	"fixed32": graphql.GqlUInt32TypeResolver,

	"uint64":  graphql.GqlUInt64TypeResolver,
	"fixed64": graphql.GqlUInt64TypeResolver,
}

type parsedFile struct {
	File          *parser.File
	Config        *SwaggerFileConfig
	OutputPath    string
	OutputPkg     string
	OutputPkgName string
}

func (g *Plugin) fileOutputPath(cfg *SwaggerFileConfig) (string, error) {
	if cfg.GetOutputPath() == "" {
		absFilePath, err := filepath.Abs(cfg.Path)
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve file absolute path")
		}
		fileName := filepath.Base(cfg.Path)
		pkg, err := GoPackageByPath(filepath.Dir(absFilePath), g.generateConfig.VendorPath)
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
	return filepath.Join(cfg.OutputPath, strings.TrimSuffix(filepath.Base(cfg.Path), ".proto")+".go"), nil
}

func (g *Plugin) fileOutputPackage(cfg *SwaggerFileConfig) (name, pkg string, err error) {
	outPath, err := g.fileOutputPath(cfg)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file output path")
	}
	pkg, err = GoPackageByPath(filepath.Dir(outPath), g.generateConfig.VendorPath)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to resolve file go package")
	}
	return strings.Replace(filepath.Base(pkg), "-", "_", -1), pkg, nil
}
