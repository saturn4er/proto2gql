package swagger2gql

import (
	"fmt"

	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/proto2gql/parser"
)

func (g *Plugin) enumTypeResolver(enumFile *parsedFile, enum *parser.Enum) (graphql.TypeResolver, error) {
	return func(ctx graphql.BodyContext) string {
		return ctx.Importer.Prefix(enumFile.OutputPkg) + g.enumVariable(enumFile, enum)
	}, nil
}

func (g *Plugin) enumGraphQLName(enumFile *parsedFile, enum *parser.Enum) string {
	return enumFile.Config.GetGQLEnumsPrefix() + snakeCamelCaseSlice(enum.TypeName)
}

func (g *Plugin) enumVariable(enumFile *parsedFile, enum *parser.Enum) string {
	return enumFile.Config.GetGQLEnumsPrefix() + snakeCamelCaseSlice(enum.TypeName)
}

func (g *Plugin) prepareFileEnums(file *parsedFile) ([]graphql.Enum, error) {
	var res []graphql.Enum
	for path, methods := range file.File.Paths {
		for method, endpoint := range methods {
			for _, parameter := range endpoint.Parameters {
				fmt.Println(parameter.)
			}
		}
		vals := make([]graphql.EnumValue, len(enum.Values))
		for i, value := range enum.Values {
			vals[i] = graphql.EnumValue{
				Name:    value.Name,
				Value:   value.Value,
				Comment: value.QuotedComment,
			}
		}
		res = append(res, graphql.Enum{
			VariableName: g.enumVariable(file, enum),
			GraphQLName:  g.enumGraphQLName(file, enum),
			Comment:      enum.QuotedComment,
			Values:       vals,
		})
	}
	return res, nil
}
