package proto

import (
	"reflect"

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
func (g *Generator) oneOfValueAssigningWrapper(msg *parser.Message, field *parser.Field) common.AssigningWrapper {
	return func(arg string, ctx common.BodyContext) string {
		pkg := g.fileGRPCSourcesPackage(msg.Type.File)
		return "&" + ctx.Importer.Prefix(pkg) + camelCaseSlice(msg.TypeName) + "_" + camelCase(field.Name) + "{" + arg + "}"
	}
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
			var fields []common.InputObjectResolverOneOfField
			for _, field := range oneOf.Fields {
				resolver, withErr, err := g.TypeValueResolver(file.File, field.Type)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get type value resolver")
				}
				fields = append(fields, common.InputObjectResolverOneOfField{
					GraphQLInputFieldName: field.Name,
					ValueResolver:         resolver,
					ResolverWithError:     withErr,
					AssigningWrapper:      g.oneOfValueAssigningWrapper(msg, field),
				})
			}
			oneOffs = append(oneOffs, common.InputObjectResolverOneOf{
				OutputFieldName: camelCase(oneOf.Name),
				Fields:          fields,
			})
		}
		var fields []common.InputObjectResolverField
		for _, field := range msg.Fields {
			resolver, withErr, err := g.TypeValueResolver(file.File, field.Type)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get type value resolver")
			}
			goType, err := goTypeByParserType(field.Type)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get go type by parser type")
			}
			if field.Repeated {
				gt := goType
				goType = common.GoType{
					Kind:     reflect.Slice,
					ElemType: &gt,
				}
			}
			fields = append(fields, common.InputObjectResolverField{
				GraphQLInputFieldName: field.Name,
				OutputFieldName:       camelCase(field.Name),
				ValueResolver:         resolver,
				ResolverWithError:     withErr,
				GoType:                goType,
			})
		}
		res = append(res, common.InputObjectResolver{
			FunctionName: g.inputMessageResolverName(msg),
			OutputGoType: goType,
			OneOfFields:  oneOffs,
			Fields:       fields,
		})

	}
	return res, nil
}
