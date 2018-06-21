package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) TypeValueResolver(parentFile *parser.File, typ *parser.Type) (_ common.ValueResolver, withErr bool, err error) {
	if typ.IsScalar() {
		gt, ok := goTypesScalars[typ.Scalar]
		if !ok {
			panic("unknown scalar: " + typ.Scalar)
		}
		return func(arg string, ctx common.BodyContext) string {
			return arg + ".(" + gt.Kind.String() + ")"
		}, false, nil
	}
	if typ.IsEnum() {
		return func(arg string, ctx common.BodyContext) string {
			return ctx.Importer.Prefix(g.fileGRPCSourcesPackage(typ.File)) + snakeCamelCaseSlice(typ.Enum.TypeName) + "(" + arg + ".(int))"
		}, false, nil
	}
	if typ.IsMessage() {
		_, pkg, err := g.fileOutputPackage(typ.File)
		if err != nil {
			return nil, false, errors.Wrapf(err, "failed to resolve type %s output package", typ)
		}
		return func(arg string, ctx common.BodyContext) string {
			return ctx.Importer.Prefix(pkg) + g.inputMessageResolverName(typ.Message) + "(tr," + arg + ")"
		}, true, nil
	}
	return func(arg string, ctx common.BodyContext) string {
		return arg + "// not implemented"
	}, false, nil

}
