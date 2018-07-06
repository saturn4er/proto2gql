package swagger2gql

import (
	"strings"

	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

func (g *Plugin) inputMessageGraphQLName(file *parsedFile, message *parser.Type) string {
	return file.Config.GetGQLMessagePrefix() + strings.Join(message.Route, "__") + "Input"
}
func (g *Plugin) inputMessageVariable(msgFile *parsedFile, message *parser.Type) string {
	return msgFile.Config.GetGQLMessagePrefix() + strings.Join(message.Route, "") + "Input"
}

//
func (g *Plugin) inputMessageTypeResolver(msgFile *parsedFile, typ *parser.Type) graphql.TypeResolver {
	if len(typ.Object.Properties) == 0 {
		return graphql.GqlNoDataTypeResolver
	}

	return func(ctx graphql.BodyContext) string {
		return ctx.Importer.Prefix(msgFile.OutputPkg) + g.inputMessageVariable(msgFile, typ)
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
func (g *Plugin) graphqlInputTypeResolver(typeFile *parsedFile, typ *parser.Type) graphql.TypeResolver {
	if typ == nil {
		return graphql.GqlNoDataTypeResolver
	}
	switch typ.Type {
	case parser.TypeObject:
		return g.inputMessageTypeResolver(typeFile, typ)
	case parser.TypeArray:
		return graphql.GqlListTypeResolver(g.graphqlInputTypeResolver(typeFile, typ.ElemType))
	case parser.TypeBoolean:
		return graphql.GqlBoolTypeResolver
	case parser.TypeFloat64:
		return graphql.GqlFloat64TypeResolver
	case parser.TypeFloat32:
		return graphql.GqlFloat32TypeResolver
	case parser.TypeInt64:
		return graphql.GqlInt64TypeResolver
	case parser.TypeInt32:
		return graphql.GqlInt32TypeResolver
	case parser.TypeString:
		return graphql.GqlStringTypeResolver
		// TODO: map
	}
	return func(ctx graphql.BodyContext) string {
		return "NOT IMPLEMENTED"
	}
}
func (g *Plugin) fileInputObjects(file *parsedFile) ([]graphql.InputObject, error) {
	var res []graphql.InputObject
	var handledObjects = map[string]struct{}{}
	var handleType func(typ *parser.Type)
	handleType = func(typ *parser.Type) {
		switch typ.Type {
		case parser.TypeObject:
			if _, handled := handledObjects[camelCaseSlice(typ.Route)]; handled {
				return
			}
			var fields []graphql.ObjectField
			for _, property := range typ.Object.Properties {
				handleType(property.Type)
				typeResolver := g.graphqlInputTypeResolver(file, property.Type)
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
				VariableName: g.inputMessageVariable(file, typ),
				GraphQLName:  g.inputMessageGraphQLName(file, typ),
				Fields:       fields,
			})

			handledObjects[camelCaseSlice(typ.Route)] = struct{}{}
		case parser.TypeArray:
			handleType(typ.ElemType)
		}
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
