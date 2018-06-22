package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) enumTypeResolver(currentFile *parser.File, enum *parser.Enum) (common.TypeResolver, error) {
	_, pkg, err := g.fileOutputPackage(enum.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file output package")
	}
	return func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(pkg) + g.enumVariable(enum)
	}, nil
}

func (g *Generator) enumGraphQLName(res *parser.Enum) string {
	cfg := g.fileConfig(res.File)
	return cfg.GetGQLEnumsPrefix() + res.Name
}

func (g *Generator) enumVariable(res *parser.Enum) string {
	return res.Name
}

func (g *Generator) prepareFileEnums(file parsedFile) ([]common.Enum, error) {
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
			VariableName: g.enumVariable(enum),
			GraphQLName:  g.enumGraphQLName(enum),
			Comment:      enum.QuotedComment,
			Values:       vals,
		})
	}
	return res, nil
}
