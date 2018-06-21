package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) inputMessageResolverName(message *parser.Message) string {
	return "Resolve" + camelCaseSlice(message.TypeName)
}
func (g *Generator) fileMessageInputObjectsResolvers(file parsedFile) ([]common.InputObjectResolver, error) {
	var res []common.InputObjectResolver
	for _, msg := range file.File.Messages {
		goType, err := goTypeByParserType(msg.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve message go type")
		}
		res = append(res, common.InputObjectResolver{
			FunctionName: g.inputMessageResolverName(msg),
			OutputGoType: goType,
		})

	}
	return res, nil
}
