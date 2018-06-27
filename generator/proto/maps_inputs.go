package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) inputMapGraphQLName(mapFile *parsedFile, res *parser.Map) string {
	return g.inputMessageVariable(mapFile, res.Message) + "__" + camelCase(res.Field.Name)
}

func (g *Generator) inputMapVariable(mapFile *parsedFile, res *parser.Map) string {
	return g.inputMessageVariable(mapFile, res.Message) + "__" + camelCase(res.Field.Name)
}

func (g *Generator) fileMapInputObjects(file *parsedFile) ([]common.MapInputObject, error) {
	var res []common.MapInputObject
	for _, msg := range file.File.Messages {
		for _, mapFld := range msg.MapFields {
			keyTypResolver, err := g.TypeOutputTypeResolver(file, mapFld.Map.KeyType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve key input type resolver")
			}
			valueFile, err := g.parsedFile(mapFld.Map.ValueType.File)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value type file")
			}
			valueTypResolver, err := g.TypeOutputTypeResolver(valueFile, mapFld.Map.ValueType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value input type resolver")
			}

			res = append(res, common.MapInputObject{
				VariableName:    g.inputMapVariable(file, mapFld.Map),
				GraphQLName:     g.inputMapGraphQLName(file, mapFld.Map),
				KeyObjectType:   keyTypResolver,
				ValueObjectType: valueTypResolver,
			})
		}

	}
	return res, nil
}
