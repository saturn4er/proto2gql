package proto

import (
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) enumTypeResolver(enumFile *parsedFile, enum *parser.Enum) (common.TypeResolver, error) {
	return func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(enumFile.OutputPkg) + g.enumVariable(enumFile, enum)
	}, nil
}

func (g *Generator) enumGraphQLName(enumFile *parsedFile, enum *parser.Enum) string {
	return enumFile.Config.GetGQLEnumsPrefix() + snakeCamelCaseSlice(enum.TypeName)
}

func (g *Generator) enumVariable(enumFile *parsedFile, enum *parser.Enum) string {
	return enumFile.Config.GetGQLEnumsPrefix() + snakeCamelCaseSlice(enum.TypeName)
}

func (g *Generator) prepareFileEnums(file *parsedFile) ([]common.Enum, error) {
	var res []common.Enum
	for _, enum := range file.File.Enums {
		vals := make([]common.EnumValue, len(enum.Values))
		for i, value := range enum.Values {
			vals[i] = common.EnumValue{
				Name:    value.Name,
				Value:   value.Value,
				Comment: value.QuotedComment,
			}
		}
		res = append(res, common.Enum{
			VariableName: g.enumVariable(file, enum),
			GraphQLName:  g.enumGraphQLName(file, enum),
			Comment:      enum.QuotedComment,
			Values:       vals,
		})
	}
	return res, nil
}
