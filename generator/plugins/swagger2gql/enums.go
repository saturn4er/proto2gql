package swagger2gql

// func (g *Plugin) enumTypeResolver(enumFile *parsedFile, enum *parser.Type) (graphql.TypeResolver, error) {
// 	return func(ctx graphql.BodyContext) string {
// 		return ctx.Importer.Prefix(enumFile.OutputPkg) + g.enumVariable(enumFile, enum)
// 	}, nil
// }
//
// func (g *Plugin) enumGraphQLName(enumFile *parsedFile, enum *parser.Type) string {
// 	return enumFile.Config.GetGQLEnumsPrefix() + snakeCamelCaseSlice(enum.Route)
// }
//
// func (g *Plugin) enumVariable(enumFile *parsedFile, enum *parser.Type) string {
// 	return enumFile.Config.GetGQLEnumsPrefix() + snakeCamelCaseSlice(enum.Route)
// }
//
// func (g *Plugin) prepareFileEnums(file *parsedFile) ([]graphql.Enum, error) {
// 	var res []graphql.Enum
// 	var handledEnums = map[string]struct{}{}
// 	var handleType func(typ *parser.Type)
// 	handleType = func(typ *parser.Type) {
// 		switch typ.Type {
// 		case parser.KindString:
// 			_, handled := handledEnums[snakeCamelCaseSlice(typ.Route)]
// 			if len(typ.Enum) == 0 || handled {
// 				return
// 			}
// 			values := make([]graphql.EnumValue, len(typ.Enum))
// 			for i, value := range typ.Enum {
// 				values[i] = graphql.EnumValue{
// 					Name:    value,
// 					Value:   i,
// 					Comment: `""`,
// 				}
// 			}
// 			res = append(res, graphql.Enum{
// 				VariableName: g.enumVariable(file, typ),
// 				GraphQLName:  g.enumGraphQLName(file, typ),
// 				Values:       values,
// 			})
// 			handledEnums[snakeCamelCaseSlice(typ.Route)] = struct{}{}
// 		case parser.KindObject:
// 			for _, property := range typ.Object.Properties {
// 				handleType(property.Type)
// 			}
// 		case parser.KindArray:
// 			handleType(typ.ElemType)
// 		}
// 	}
// 	for _, tag := range file.File.Tags {
// 		for _, method := range tag.Methods {
// 			for _, param := range method.Parameters {
// 				handleType(param.Type)
// 			}
// 			for _, response := range method.Responses {
// 				handleType(response.ResultType)
// 			}
// 		}
// 	}
// 	return res, nil
// }
