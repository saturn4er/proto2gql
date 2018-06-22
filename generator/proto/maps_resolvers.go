package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) mapResolverFunctionName(mp *parser.Map) string {
	return "Resolve" + g.inputMapVariable(mp)
}
func (g *Generator) fileInputMapResolvers(file parsedFile) ([]common.MapInputObjectResolver, error) {
	var res []common.MapInputObjectResolver
	for _, msg := range file.File.Messages {
		for _, mapFld := range msg.MapFields {
			keyGoType, err := goTypeByParserType(mapFld.Map.KeyType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve key go type")
			}
			valueGoType, err := goTypeByParserType(mapFld.Map.ValueType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value go type")
			}
			keyTypeResolver, keyResolveWithErr, err := g.TypeValueResolver(file.File, mapFld.Map.KeyType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get key type value resolver")
			}
			valueTypeResolver, valueResolveWithErr, err := g.TypeValueResolver(file.File, mapFld.Map.ValueType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get value type value resolver")
			}
			res = append(res, common.MapInputObjectResolver{
				FunctionName:           g.mapResolverFunctionName(mapFld.Map),
				KeyGoType:              keyGoType,
				ValueGoType:            valueGoType,
				KeyResolver:            keyTypeResolver,
				KeyResolverWithError:   keyResolveWithErr,
				ValueResolver:          valueTypeResolver,
				ValueResolverWithError: valueResolveWithErr,
			})
		}

	}
	return res, nil
}
