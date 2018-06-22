package proto

import (
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) inputMessageResolverName(message *parser.Message) string {
	return "Resolve" + camelCaseSlice(message.TypeName)
}
func (g *Generator) oneOfValueResolver(oneof *parser.OneOf, field *parser.Field) (_ common.ValueResolver, withErr bool) {
	return func(arg string, ctx common.BodyContext) string {
		return arg
	}, true
}
func (g *Generator) fileMessageInputObjectsResolvers(file parsedFile) ([]common.InputObjectResolver, error) {
	var res []common.InputObjectResolver
	for _, msg := range file.File.Messages {
		goType, err := goTypeByParserType(msg.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve message go type")
		}
		var oneOffs []common.InputObjectResolverOneOf
		for _, oneOf := range msg.OneOffs {
			var fields []common.InputObjectResolverField
			for _, field := range oneOf.Fields {
				resolver, withErr := g.oneOfValueResolver(oneOf, field)
				fields = append(fields, common.InputObjectResolverField{
					GraphqlInputField: field.Name,
					ValueResolver:     resolver,
					ResolverWithError: withErr,
				})
			}
			oneOffs = append(oneOffs, common.InputObjectResolverOneOf{
				OutputFieldName: camelCase(oneOf.Name),
				Fields:          fields,
			})
		}
		res = append(res, common.InputObjectResolver{
			FunctionName: g.inputMessageResolverName(msg),
			OutputGoType: goType,
			OneOfFields:  oneOffs,
		})

	}
	return res, nil
}
