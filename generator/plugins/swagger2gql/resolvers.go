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
	parser.KindNull:    graphql.GqlNoDataTypeResolver,
	parser.KindFile:    graphql.GqlFileTypeResolver,
}

func (p *Plugin) TypeOutputTypeResolver(typeFile *parsedFile, typ parser.Type, required bool) (graphql.TypeResolver, error) {
	var res graphql.TypeResolver
	switch t := typ.(type) {
	case parser.Scalar:
		resolver, ok := scalarsResolvers[typ.Kind()]
		if !ok {
			return nil, errors.Errorf(": %s", typ.Kind())
		}
		res = resolver
	case parser.Object:
		msgResolver, err := p.outputMessageTypeResolver(typeFile, t)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get message type resolver")
		}
		res = msgResolver
	case parser.Array:
		elemResolver, err := p.TypeOutputTypeResolver(typeFile, t.ElemType, true)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get array element type resolver")
		}
		res = graphql.GqlListTypeResolver(elemResolver)
	case parser.Map:
		res = func(ctx graphql.BodyContext) string {
			return p.mapOutputObjectVariable(typeFile, t)
		}
		res = graphql.GqlListTypeResolver(graphql.GqlNonNullTypeResolver(res))
	default:
		return nil, errors.Errorf("not implemented %v", typ.Kind())
	}
	if required {
		res = graphql.GqlNonNullTypeResolver(res)
	}
	return res, nil
}
func (p *Plugin) TypeInputTypeResolver(typeFile *parsedFile, typ parser.Type) (graphql.TypeResolver, error) {
	switch t := typ.(type) {
	case parser.Scalar:
		resolver, ok := scalarsResolvers[t.Kind()]
		if !ok {
			return nil, errors.Errorf("unimplemented scalar type: %s", t.Kind())
		}
		return resolver, nil
	case parser.Object:
		return p.inputObjectTypeResolver(typeFile, t), nil
	case parser.Array:
		elemResolver, err := p.TypeOutputTypeResolver(typeFile, t.ElemType, true)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get array element type resolver")
		}
		return graphql.GqlListTypeResolver(elemResolver), nil
	}
	return nil, errors.New("not implemented " + typ.Kind().String())
}
func (p *Plugin) TypeValueResolver(file *parsedFile, typ parser.Type, required bool, ctxKey string) (_ graphql.ValueResolver, withErr bool, err error) {
	if ctxKey != "" {
		goType, err := p.goTypeByParserType(file, typ, true)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed to resolve go type")
		}
		return func(arg string, ctx graphql.BodyContext) string {
			return `ctx.Value("` + ctxKey + `").(` + goType.String(ctx.Importer) + `)`
		}, false, nil
	}
	goTyp, err := p.goTypeByParserType(file, typ, true)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to resolve go type by parser type")
	}
	switch t := typ.(type) {
	case parser.Scalar:
		if t.Kind() == parser.KindFile{
			return func(arg string, ctx graphql.BodyContext) string {
					return `ctx.Value("files")`
			}, false, nil
		}
		goTyp, ok := scalarsGoTypesNames[typ.Kind()]
		if !ok {
			return nil, false, errors.Errorf("scalar %s is not implemented", typ.Kind())
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
		return graphql.ResolverCall(file.OutputPkg, "Resolve"+snakeCamelCaseSlice(t.Route)), true, nil

	case parser.Array:
		elemResolver, elemResolverWithErr, err := p.TypeValueResolver(file, t.ElemType, false, "")
		if err != nil {
			return nil, false, errors.Wrap(err, "failed to get array element type value resolver")
		}

		return func(arg string, ctx graphql.BodyContext) string {
			res, err := p.renderArrayValueResolver(arg, goTyp, ctx, elemResolver, elemResolverWithErr)
			if err != nil {
				panic(err)
			}
			return res
		}, true, nil
	case parser.Map:
		return func(arg string, ctx graphql.BodyContext) string {
			if ctx.TracerEnabled {
				return "Resolve" + p.mapInputObjectVariable(file, t) + "(tr, tr.ContextWithSpan(ctx, span), " + arg + ")"
			}
			return "Resolve" + p.mapInputObjectVariable(file, t) + "(ctx, " + arg + ")"
		}, true, nil
	}
	return nil, false, errors.Errorf("unknown type: %v", typ.Kind().String())

}