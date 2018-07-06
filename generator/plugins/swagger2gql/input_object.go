package swagger2gql

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

func (g *Plugin) inputObjectGQLName(file *parsedFile, obj parser.Object) string {
	return file.Config.GetGQLMessagePrefix() + strings.Join(obj.Route, "__") + "Input"
}
func (g *Plugin) inputObjectVariable(msgFile *parsedFile, obj parser.Object) string {
	return msgFile.Config.GetGQLMessagePrefix() + strings.Join(obj.Route, "") + "Input"
}

//
func (g *Plugin) inputObjectTypeResolver(msgFile *parsedFile, obj parser.Object) graphql.TypeResolver {
	if len(obj.Properties) == 0 {
		return graphql.GqlNoDataTypeResolver
	}

	return func(ctx graphql.BodyContext) string {
		return ctx.Importer.Prefix(msgFile.OutputPkg) + g.inputObjectVariable(msgFile, obj)
	}
}

//
// func (g *Plugin) inputMessageFieldTypeResolver(file *parsedFile, field *parser.Field) (graphql.TypeResolver, error) {
// 	resolver, err := g.TypeOutputTypeResolver(file, field.Type)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to get input type resolver")
// 	}
// 	if field.Repeated {
// 		resolver = graphql.GqlListTypeResolver(graphql.GqlNonNullTypeResolver(resolver))
// 	}
// 	return resolver, nil
// }
//
// func (g *Plugin) outputObjectMapFieldTypeResolver(mapFile *parsedFile, mp *parser.Map) (graphql.TypeResolver, error) {
// 	res := func(ctx graphql.BodyContext) string {
// 		return ctx.Importer.Prefix(mapFile.OutputPkg) + g.outputMapVariable(mapFile, mp)
// 	}
// 	return graphql.GqlListTypeResolver(graphql.GqlNonNullTypeResolver(res)), nil
// }
// func (g *Plugin) inputObjectMapFieldTypeResolver(mapFile *parsedFile, mp *parser.Map) (graphql.TypeResolver, error) {
// 	res := func(ctx graphql.BodyContext) string {
// 		return ctx.Importer.Prefix(mapFile.OutputPkg) + g.inputMapVariable(mapFile, mp)
// 	}
// 	return graphql.GqlListTypeResolver(graphql.GqlNonNullTypeResolver(res)), nil
// }

func (g *Plugin) fileInputObjects(file *parsedFile) ([]graphql.InputObject, error) {
	var res []graphql.InputObject
	var handledObjects = map[string]struct{}{}
	var handleType func(typ parser.Type) error
	handleType = func(typ parser.Type) error {
		switch t := typ.(type) {
		case parser.Object:
			if _, handled := handledObjects[camelCaseSlice(t.Route)]; handled {
				return nil
			}
			var fields []graphql.ObjectField
			for _, property := range t.Properties {
				if err := handleType(property.Type); err != nil {
					return err
				}
				typeResolver, err := g.TypeInputTypeResolver(file, property.Type)
				if err != nil {
					return errors.Wrap(err, "failed to get input type resolver")
				}
				if property.Required {
					typeResolver = graphql.GqlNonNullTypeResolver(typeResolver)
				}
				fields = append(fields, graphql.ObjectField{
					Name:           property.Name,
					Type:           typeResolver,
					GoObjectGetter: "",
					NeedCast:       false,
				})
			}
			res = append(res, graphql.InputObject{
				VariableName: g.inputObjectVariable(file, t),
				GraphQLName:  g.inputObjectGQLName(file, t),
				Fields:       fields,
			})

			handledObjects[camelCaseSlice(t.Route)] = struct{}{}
		case parser.Array:
			return handleType(t.ElemType)
		}
		return nil
	}
	for _, tag := range file.File.Tags {
		for _, method := range tag.Methods {
			for _, parameter := range method.Parameters {
				handleType(parameter.Type)
			}
		}
	}
	return res, nil
}
