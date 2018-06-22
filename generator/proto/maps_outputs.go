package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) outputMapGraphQLName(res *parser.Map) string {
	return g.outputMessageVariable(res.Message) + "__" + res.Field.Name
}

func (g *Generator) outputMapVariable(res *parser.Map) string {
	return g.outputMessageVariable(res.Message) + "__" + res.Field.Name
}

func (g *Generator) fileMapOutputObjects(file parsedFile) ([]common.MapOutputObject, error) {
	var res []common.MapOutputObject
	for _, msg := range file.File.Messages {
		for _, mapFld := range msg.MapFields {
			keyTypResolver, err := g.TypeOutputTypeResolver(file.File, mapFld.Map.KeyType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve key input type resolver")
			}
			valueTypResolver, err := g.TypeOutputTypeResolver(file.File, mapFld.Map.ValueType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value input type resolver")
			}
			res = append(res, common.MapOutputObject{
				VariableName:    g.outputMapVariable(mapFld.Map),
				GraphQLName:     g.outputMapGraphQLName(mapFld.Map),
				KeyObjectType:   keyTypResolver,
				ValueObjectType: valueTypResolver,
			})
		}
	}
	return res, nil
}
