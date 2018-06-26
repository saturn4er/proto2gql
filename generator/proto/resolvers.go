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
		file, err := g.parsedFile(typ.File)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve type parsed file")
		}
		res, err := g.enumTypeResolver(file, typ.Enum)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get enum type resolver")
		}
		return res, nil
	}
	return nil, errors.New("not implemented " + typ.String())
}
func (g *Generator) TypeInputTypeResolver(typeFile *parsedFile, typ *parser.Type) (common.TypeResolver, error) {
	if typ.IsScalar() {
		resolver, ok := scalarsResolvers[typ.Scalar]
		if !ok {
			return nil, errors.Errorf("unimplemented scalar type: %s", typ.Scalar)
		}
		return resolver, nil
	}
	if typ.IsMessage() {
		res, err := g.inputMessageTypeResolver(typeFile, typ.Message)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get message type resolver")
		}
		return res, nil
	}
	if typ.IsEnum() {
		res, err := g.enumTypeResolver(typeFile, typ.Enum)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get enum type resolver")
		}
		return res, nil
	}
	if typ.IsMap() {
		return g.inputObjectMapFieldTypeResolver(typeFile, typ.Map)
	}
	return nil, errors.New("not implemented " + typ.String())
}
func (g *Generator) TypeValueResolver(typeFile *parsedFile, typ *parser.Type, ctxKey string) (_ common.ValueResolver, withErr bool, err error) {
	if ctxKey != "" {
		goType, err := g.goTypeByParserType(typ)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed to resolve go type")
		}
		return func(arg string, ctx common.BodyContext) string {
			return `ctx.Value("` + ctxKey + `").(` + goType.String(ctx.Importer) + `)`
		}, false, nil
	}
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
			return ctx.Importer.Prefix(typeFile.GRPCSourcesPkg) + snakeCamelCaseSlice(typ.Enum.TypeName) + "(" + arg + ".(int))"
		}, false, nil
	}
	if typ.IsMessage() {
		return func(arg string, ctx common.BodyContext) string {
			if ctx.TracerEnabled {
				return ctx.Importer.Prefix(typeFile.OutputPkg) + g.inputMessageResolverName(typeFile, typ.Message) + "(tr, tr.ContextWithSpan(ctx,span), " + arg + ")"
			} else {
				return ctx.Importer.Prefix(typeFile.OutputPkg) + g.inputMessageResolverName(typeFile, typ.Message) + "(ctx, " + arg + ")"
			}
		}, true, nil
	}
	if typ.IsMap() {
		return func(arg string, ctx common.BodyContext) string {
			if ctx.TracerEnabled {
				return ctx.Importer.Prefix(typeFile.OutputPkg) + g.mapResolverFunctionName(typeFile, typ.Map) + "(tr, tr.ContextWithSpan(ctx,span), " + arg + ")"
			} else {
				return ctx.Importer.Prefix(typeFile.OutputPkg) + g.mapResolverFunctionName(typeFile, typ.Map) + "(ctx, " + arg + ")"
			}
		}, true, nil
	}
	return func(arg string, ctx common.BodyContext) string {
		return arg + "// not implemented"
	}, false, nil

}
