package swagger2gql

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

var scalarsResolvers = map[parser.Kind]graphql.TypeResolver{
	parser.KindBoolean: graphql.GqlBoolTypeResolver,
	parser.KindFloat64: graphql.GqlFloat64TypeResolver,
	parser.KindFloat32: graphql.GqlFloat32TypeResolver,
	parser.KindInt64:   graphql.GqlInt64TypeResolver,
	parser.KindInt32:   graphql.GqlInt32TypeResolver,
	parser.KindString:  graphql.GqlStringTypeResolver,
}

func (g *Plugin) TypeOutputTypeResolver(typeFile *parsedFile, typ parser.Type) (graphql.TypeResolver, error) {
	switch t := typ.(type) {
	case parser.Scalar:
		resolver, ok := scalarsResolvers[typ.Kind()]
		if !ok {
			return nil, errors.Errorf(": %s", typ.Kind())
		}
		return resolver, nil
	case parser.Object:
		res, err := g.outputMessageTypeResolver(typeFile, t)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get message type resolver")
		}
		return res, nil
	case parser.Array:
		elemResolver, err := g.TypeOutputTypeResolver(typeFile, t.ElemType)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get array element type resolver")
		}
		return graphql.GqlListTypeResolver(elemResolver), nil
	case parser.Map:
		return func(ctx graphql.BodyContext) string {
			return "not implementred"
		}, nil
	}
	return nil, errors.Errorf("not implemented %v", typ.Kind())
}
func (g *Plugin) TypeInputTypeResolver(typeFile *parsedFile, typ parser.Type) (graphql.TypeResolver, error) {
	switch t := typ.(type) {
	case parser.Scalar:
		resolver, ok := scalarsResolvers[t.Kind()]
		if !ok {
			return nil, errors.Errorf("unimplemented scalar type: %s", t.Kind())
		}
		return resolver, nil
	case parser.Object:
		return g.inputObjectTypeResolver(typeFile, t), nil
	case parser.Array:
		elemResolver, err := g.TypeOutputTypeResolver(typeFile, t.ElemType)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get array element type resolver")
		}
		return graphql.GqlListTypeResolver(elemResolver), nil
	}
	return nil, errors.New("not implemented " + typ.Kind().String())
}
func (g *Plugin) TypeValueResolver(file *parsedFile, typ parser.Type, required bool, ctxKey string) (_ graphql.ValueResolver, withErr bool, err error) {
	goTyp, err := g.goTypeByParserType(file, typ, true)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to resolve go type by parser type")
	}
	switch t := typ.(type) {
	case parser.Scalar:
		goTyp, ok := scalarsGoTypesNames[typ.Kind()]
		if !ok {
			return nil, false, errors.Wrapf(err, "scalar %s is not implemented", typ.Kind())
		}
		return func(arg string, ctx graphql.BodyContext) string {

			if !required {
				return arg + ".(" + goTyp + ")"
			}
			return "func(arg interface{}) *" + goTyp + "{\n" +
				"val := arg.(" + goTyp + ")\n" +
				"return &val\n" +
				"}(" + arg + ")"
		}, false, nil
	case parser.Object:
		return func(arg string, ctx graphql.BodyContext) string {
			if ctx.TracerEnabled {
				return "Resolve" + snakeCamelCaseSlice(t.Route) + "(tr, tr.ContextWithSpan(ctx, span), " + arg + ")"
			}
			return "Resolve" + snakeCamelCaseSlice(t.Route) + "(ctx, " + arg + ")"
		}, true, nil
	case parser.Array:
		elemResolver, elemResolverWithErr, err := g.TypeValueResolver(file, t.ElemType, false, "")
		if err != nil {
			return nil, false, errors.Wrap(err, "failed to get array element type value resolver")
		}

		return func(arg string, ctx graphql.BodyContext) string {
			res, err := g.renderArrayValueResolver(arg, goTyp, ctx, elemResolver, elemResolverWithErr)
			if err != nil {
				panic(err)
			}
			return res
		}, true, nil
	}
	return nil, false, errors.Errorf("unknown type: %v", typ.Kind().String())

}
