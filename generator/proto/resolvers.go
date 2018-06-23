package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) TypeOutputTypeResolver(typ *parser.Type) (common.TypeResolver, error) {
	if typ.IsScalar() {
		resolver, ok := scalarsResolvers[typ.Scalar]
		if !ok {
			return nil, errors.Errorf("unimplemented scalar type: %s", typ.Scalar)
		}
		return resolver, nil
	}
	if typ.IsMessage() {
		res, err := g.outputMessageTypeResolver(typ.Message)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get message type resolver")
		}
		return res, nil
	}
	if typ.IsEnum() {
		res, err := g.enumTypeResolver(typ.Enum)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get enum type resolver")
		}
		return res, nil
	}
	return nil, errors.New("not implemented " + typ.String())
}
func (g *Generator) TypeInputTypeResolver(currentFile *parser.File, typ *parser.Type) (common.TypeResolver, error) {
	if typ.IsScalar() {
		resolver, ok := scalarsResolvers[typ.Scalar]
		if !ok {
			return nil, errors.Errorf("unimplemented scalar type: %s", typ.Scalar)
		}
		return resolver, nil
	}
	if typ.IsMessage() {
		res, err := g.inputMessageTypeResolver(currentFile, typ.Message)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get message type resolver")
		}
		return res, nil
	}
	if typ.IsEnum() {
		res, err := g.enumTypeResolver(typ.Enum)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get enum type resolver")
		}
		return res, nil
	}
	if typ.IsMap(){
		return g.inputObjectMapFieldTypeResolver(typ.Map)
	}
	return nil, errors.New("not implemented " + typ.String())
}
func (g *Generator) TypeValueResolver(typ *parser.Type) (_ common.ValueResolver, withErr bool, err error) {
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
			if ctx.TracerEnabled {
				return ctx.Importer.Prefix(pkg) + g.inputMessageResolverName(typ.Message) + "(tr, tr.ContextWithSpan(ctx,span), " + arg + ")"
			} else {
				return ctx.Importer.Prefix(pkg) + g.inputMessageResolverName(typ.Message) + "(ctx, " + arg + ")"
			}
		}, true, nil
	}
	return func(arg string, ctx common.BodyContext) string {
		return arg + "// not implemented"
	}, false, nil

}
