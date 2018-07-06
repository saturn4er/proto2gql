package swagger2gql

import (
	"bytes"
	"reflect"
	"text/template"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

// func (g *Plugin) inputMessageResolverName(msgFile *parsedFile, message *parser.Message) string {
// 	return "Resolve" + g.inputObjectGQLName(msgFile, message)
// }
//
// func (g *Plugin) oneOfValueAssigningWrapper(file *parsedFile, msg *parser.Message, field *parser.Field) graphql.AssigningWrapper {
// 	return func(arg string, ctx graphql.BodyContext) string {
// 		return "&" + ctx.Importer.Prefix(file.GRPCSourcesPkg) + camelCaseSlice(msg.TypeName) + "_" + camelCase(field.Name) + "{" + arg + "}"
// 	}
// }
var arrayValueTemplate *template.Template

func init() {
	tplBody, err := templatesArray_value_resolverGohtmlBytes()
	if err != nil {
		panic(errors.Wrap(err, "failed to get array value resolver template").Error())
	}
	tpl, err := template.New("array_value_resolver").Parse(string(tplBody))
	if err != nil {
		panic(errors.Wrap(err, "failed to parse array value resolver template"))
	}
	arrayValueTemplate = tpl
}

func (g *Plugin) renderArrayValueResolver(arg string, resultGoTyp graphql.GoType, ctx graphql.BodyContext, elemResolver graphql.ValueResolver, elemResolverWithErr bool) (string, error) {
	res := new(bytes.Buffer)
	err := arrayValueTemplate.Execute(res, map[string]interface{}{
		"resultType": func() string {
			return resultGoTyp.String(ctx.Importer)
		},
		"rootCtx":             ctx,
		"elemResolver":        elemResolver,
		"elemResolverWithErr": elemResolverWithErr,
		"arg":                 arg,
		"errorsPkg": func() string {
			return ctx.Importer.New(graphql.ErrorsPkgPath)
		},
	})
	return res.String(), err
}
func (g *Plugin) fileInputMessagesResolvers(file *parsedFile) ([]graphql.InputObjectResolver, error) {
	var res []graphql.InputObjectResolver
	var handledObjects = map[string]struct{}{}
	var handleType func(typ parser.Type) error
	handleType = func(typ parser.Type) error {
		switch t := typ.(type) {
		case parser.Array:
			return handleType(t.ElemType)
		case parser.Object:
			var fields []graphql.InputObjectResolverField
			if _, handled := handledObjects[snakeCamelCaseSlice(t.Route)]; handled {
				return nil
			}
			for _, property := range t.Properties {
				err := handleType(property.Type)
				if err != nil {
					return errors.Wrapf(err, "failed to resolve property %s objects resolvers", property.Name)
				}
				valueResolver, withErr, err := g.TypeValueResolver(file, property.Type, property.Required, "")
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
			resGoType, err := g.goTypeByParserType(file, t, true)
			if err != nil {
				return errors.Wrap(err, "failed to resolve object go type")
			}
			res = append(res, graphql.InputObjectResolver{
				FunctionName: "Resolve" + snakeCamelCaseSlice(t.Route),
				Fields:       fields,
				OutputGoType: resGoType,
			})
			handledObjects[snakeCamelCaseSlice(t.Route)] = struct{}{}
		}
		return nil
	}
	for _, tag := range file.File.Tags {
		for _, method := range tag.Methods {
			for _, parameter := range method.Parameters {
				err := handleType(parameter.Type)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to handle type %v", parameter.Type.Kind())
				}
			}
		}
	}
	return res, nil
}
