package proto2gql

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/proto2gql/parser"
)

func (g *Proto2GraphQL) TypeOutputTypeResolver(typeFile *parsedFile, typ *parser.Type) (graphql.TypeResolver, error) {
	if typ.IsScalar() {
		resolver, ok := scalarsResolvers[typ.Scalar]
		if !ok {
			return nil, errors.Errorf("unimplemented scalar type: %s", typ.Scalar)
		}
		return resolver, nil
	}
	if typ.IsMessage() {
		msgCfg, err := typeFile.Config.MessageConfig(typ.Message.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve message %s config", typ.Message.Name)
		}
		if !typ.Message.HaveFieldsExcept(msgCfg.ErrorField){
			return graphql.GqlNoDataTypeResolver, nil
		}
		res, err := g.outputMessageTypeResolver(typeFile, typ.Message)
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
	if typ.IsMap() {
		return g.outputObjectMapFieldTypeResolver(typeFile, typ.Map)
	}
	return nil, errors.New("not implemented " + typ.String())
}
func (g *Proto2GraphQL) TypeInputTypeResolver(typeFile *parsedFile, typ *parser.Type) (graphql.TypeResolver, error) {
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
func (g *Proto2GraphQL) TypeValueResolver(typeFile *parsedFile, typ *parser.Type, ctxKey string) (_ graphql.ValueResolver, withErr, fromArgs bool, err error) {
	if ctxKey != "" {
		goType, err := g.goTypeByParserType(typ)
		if err != nil {
			return nil, false, false, errors.Wrap(err, "failed to resolve go type")
		}
		return func(arg string, ctx graphql.BodyContext) string {
			return `ctx.Value("` + ctxKey + `").(` + goType.String(ctx.Importer) + `)`
		}, false, false, nil
	}
	if typ.IsScalar() {
		gt, ok := goTypesScalars[typ.Scalar]
		if !ok {
			panic("unknown scalar: " + typ.Scalar)
		}
		return func(arg string, ctx graphql.BodyContext) string {
			return arg + ".(" + gt.Kind.String() + ")"
		}, false, true, nil
	}
	if typ.IsEnum() {
		return func(arg string, ctx graphql.BodyContext) string {
			return ctx.Importer.Prefix(typeFile.GRPCSourcesPkg) + snakeCamelCaseSlice(typ.Enum.TypeName) + "(" + arg + ".(int))"
		}, false, true, nil
	}
	if typ.IsMessage() {
		return func(arg string, ctx graphql.BodyContext) string {
			if ctx.TracerEnabled {
				return ctx.Importer.Prefix(typeFile.OutputPkg) + g.inputMessageResolverName(typeFile, typ.Message) + "(tr, tr.ContextWithSpan(ctx,span), " + arg + ")"
			} else {
				return ctx.Importer.Prefix(typeFile.OutputPkg) + g.inputMessageResolverName(typeFile, typ.Message) + "(ctx, " + arg + ")"
			}
		}, true, true, nil
	}
	if typ.IsMap() {
		return graphql.ResolverCall(typeFile.OutputPkg, g.mapResolverFunctionName(typeFile, typ.Map)), true, true, nil
	}
	return func(arg string, ctx graphql.BodyContext) string {
		return arg + "// not implemented"
	}, false, true, nil

}

func (g *Proto2GraphQL) FieldOutputValueResolver(fieldFile *parsedFile, fieldName string, fieldRepeated bool, fieldType *parser.Type) (_ graphql.ValueResolver, err error) {
	switch {
	case fieldType.IsScalar(), fieldType.IsMessage():
		return graphql.IdentAccessValueResolver(camelCase(fieldName)), nil
	case fieldType.IsMap():
		goKeyTyp, err := g.goTypeByParserType(fieldType.Map.KeyType)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve field key go type")
		}
		goValueTyp, err := g.goTypeByParserType(fieldType.Map.ValueType)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve field value go type")
		}
		return func(arg string, ctx graphql.BodyContext) string {
			return "func(arg map[" + goKeyTyp.String(ctx.Importer) + "]" + goValueTyp.String(ctx.Importer) + ") []map[string]interface{} {" +
				"\n  	res := make([]int, len(arg))" +
				"\n 	for i, val := range arg {" +
				"\n 		res[i] = int(val)" +
				"\n		}" +
				"\n 	return res" +
				"\n	}(" + arg + ".Get" + camelCase(fieldName) + "())"
		}, nil
	case fieldType.IsEnum():
		if fieldRepeated {
			goTyp, err := g.goTypeByParserType(fieldType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve field go type")
			}
			return func(arg string, ctx graphql.BodyContext) string {

				return "func(arg []" + goTyp.String(ctx.Importer) + ") []int {" +
					"\n  	res := make([]int, len(arg))" +
					"\n 	for i, val := range arg {" +
					"\n 		res[i] = int(val)" +
					"\n		}" +
					"\n 	return res" +
					"\n	}(" + arg + ".Get" + camelCase(fieldName) + "())"
			}, nil
		} else {
			return func(arg string, ctx graphql.BodyContext) string {
				return "int(" + arg + ".Get" + camelCase(fieldName) + "())"
			}, nil
		}
	}
	return func(arg string, ctx graphql.BodyContext) string {
		return arg + "// not implemented"
	}, nil
}
