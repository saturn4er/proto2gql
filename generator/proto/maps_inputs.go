package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) inputMapGraphQLName(res *parser.Map) string {
	return g.inputMessageVariable(res.Message) + "__" + camelCase(res.Field.Name)
}

func (g *Generator) inputMapVariable(res *parser.Map) string {
	return g.inputMessageVariable(res.Message) + "__" + camelCase(res.Field.Name)
}

func (g *Generator) fileMapInputObjects(file *parser.File) ([]common.MapInputObject, error) {
	var res []common.MapInputObject
	for _, msg := range file.Messages {
		for _, mapFld := range msg.MapFields {
			keyTypResolver, err := g.TypeOutputTypeResolver(mapFld.Map.KeyType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve key input type resolver")
			}
			valueTypResolver, err := g.TypeOutputTypeResolver(mapFld.Map.ValueType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value input type resolver")
			}

			res = append(res, common.MapInputObject{
				VariableName:    g.inputMapVariable(mapFld.Map),
				GraphQLName:     g.inputMapGraphQLName(mapFld.Map),
				KeyObjectType:   keyTypResolver,
				ValueObjectType: valueTypResolver,
			})
		}

	}
	return res, nil
}
