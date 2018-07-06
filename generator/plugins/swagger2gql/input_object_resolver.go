package swagger2gql

import (
	"reflect"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

// func (g *Plugin) inputMessageResolverName(msgFile *parsedFile, message *parser.Message) string {
// 	return "Resolve" + g.inputMessageGraphQLName(msgFile, message)
// }
//
// func (g *Plugin) oneOfValueAssigningWrapper(file *parsedFile, msg *parser.Message, field *parser.Field) graphql.AssigningWrapper {
// 	return func(arg string, ctx graphql.BodyContext) string {
// 		return "&" + ctx.Importer.Prefix(file.GRPCSourcesPkg) + camelCaseSlice(msg.TypeName) + "_" + camelCase(field.Name) + "{" + arg + "}"
// 	}
// }
var scalarsGoTypes = map[parser.Typ]string{
	parser.TypeString:  "string",
	parser.TypeFloat32: "float32",
	parser.TypeFloat64: "float65",
	parser.TypeInt64:   "int64",
	parser.TypeInt32:   "int32",
	parser.TypeBoolean: "bool",
}

func (g *Plugin) typeValueResolver(file *parsedFile, typ *parser.Type) (graphql.ValueResolver, bool, error) {
	goTyp, err := g.goTypeByParserType(file, typ)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to resolve go type by parser type")
	}
	if goTyp, ok := scalarsGoTypes[typ.Type]; ok {
		return func(arg string, ctx graphql.BodyContext) string {
			return arg + ".(" + goTyp + ")"
		}, false, nil
	}
	switch typ.Type {
	case parser.TypeObject:
		return func(arg string, ctx graphql.BodyContext) string {
			return arg
		}, false, nil
	case parser.TypeArray:
		elemResolver, withErr, err := g.typeValueResolver(file, typ.ElemType)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed to get array element type value resolver")
		}
		return func(arg string, ctx graphql.BodyContext) string {

			res := "func (arg interface{}) (_ " + goTyp.String(ctx.Importer) + ", err error) {\n" +
				"elements, ok := arg.([]interface{})\n" +
				"if !ok {" +
				"	err = " + ctx.Importer.Prefix(graphql.ErrorsPkgPath) + `New("argument is not array")` +
				"	return" +
				"}" +
				"res := make(" + goTyp.String(ctx.Importer) + ",len(elements))\n" +
				"for i, element := range elements{\n"
			if withErr {
				res += "elVal, err := " + elemResolver("element", ctx) + "\n" +
					"if err != nil{\n" +
					"	err = " + ctx.Importer.Prefix(graphql.ErrorsPkgPath) + `Wrap(err, "can't resolve array element")` +
					"	return\n" +
					"}\n" +
					"res[i] = elVal\n"
			} else {
				res += "res[i] = " + elemResolver("element", ctx) + "\n"
			}

			res += "}\n" +
				"return res,  nil\n" +
				"\n}(" + arg + ")"
			return res
		}, true, nil
	}
	return nil, false, errors.Errorf("unknown type: %v", typ)
}
func (g *Plugin) fileInputMessagesResolvers(file *parsedFile) ([]graphql.InputObjectResolver, error) {
	var res []graphql.InputObjectResolver
	// var handledObjects = map[string]struct{}{}
	var handleType func(typ *parser.Type) error
	handleType = func(typ *parser.Type) error {
		switch typ.Type {
		case parser.TypeNull:
			return nil
		case parser.TypeObject:
			var fields []graphql.InputObjectResolverField
			for _, property := range typ.Object.Properties {
				valueResolver, withErr, err := g.typeValueResolver(file, property.Type)
				if err != nil {
					return errors.Wrap(err, "failed to get property value resolver")
				}
				fields = append(fields, graphql.InputObjectResolverField{
					GraphQLInputFieldName: property.Name,
					OutputFieldName:       pascalize(camelCase(property.Name)),
					ValueResolver:         valueResolver,
					ResolverWithError:     withErr,
					GoType: graphql.GoType{
						Kind:   reflect.Uint,
						Scalar: true,
					},
				})
			}
			resGoType, err := g.goTypeByParserType(file, typ)
			if err != nil {
				return errors.Wrap(err, "failed to resolve object go type")
			}
			res = append(res, graphql.InputObjectResolver{
				FunctionName: "Resolve" + snakeCamelCaseSlice(typ.Route),
				Fields:       fields,
				OutputGoType: resGoType,
			})
			return nil

		}
		return errors.New("unknown type")
	}
	for _, tag := range file.File.Tags {
		for _, method := range tag.Methods {
			for _, parameter := range method.Parameters {
				err := handleType(parameter.Type)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to handle type %v", parameter.Type)
				}
			}
		}
	}

	return res, nil
}
