package proto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) inputMessageResolverName(msgFile *parsedFile, message *parser.Message) string {
	return "Resolve" + g.inputMessageGraphQLName(msgFile, message)
}

func (g *Generator) oneOfValueAssigningWrapper(file *parsedFile, msg *parser.Message, field *parser.Field) common.AssigningWrapper {
	return func(arg string, ctx common.BodyContext) string {
		return "&" + ctx.Importer.Prefix(file.GRPCSourcesPkg) + camelCaseSlice(msg.TypeName) + "_" + camelCase(field.Name) + "{" + arg + "}"
	}
}

func (g *Generator) fileInputMessagesResolvers(file *parsedFile) ([]common.InputObjectResolver, error) {
	var res []common.InputObjectResolver
	for _, msg := range file.File.Messages {
		msgCfg, err := file.Config.MessageConfig(dotedTypeName(msg.TypeName))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve message '%s' config", dotedTypeName(msg.TypeName))
		}
		var oneOffs []common.InputObjectResolverOneOf
		for _, oneOf := range msg.OneOffs {
			var fields []common.InputObjectResolverOneOfField
			for _, fld := range oneOf.Fields {
				fldTypeFile, err := g.parsedFile(fld.Type.File)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to resolve message '%s' field '%s' type parsed file", dotedTypeName(msg.TypeName), fld)
				}
				fldCfg := msgCfg.Fields[fld.Name]
				resolver, withErr, err := g.TypeValueResolver(fldTypeFile, fld.Type, fldCfg.ContextKey)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get type value resolver")
				}
				fields = append(fields, common.InputObjectResolverOneOfField{
					GraphQLInputFieldName: fld.Name,
					ValueResolver:         resolver,
					ResolverWithError:     withErr,
					AssigningWrapper:      g.oneOfValueAssigningWrapper(file, msg, fld),
				})
			}
			oneOffs = append(oneOffs, common.InputObjectResolverOneOf{
				OutputFieldName: camelCase(oneOf.Name),
				Fields:          fields,
			})
		}
		var fields []common.InputObjectResolverField
		for _, fld := range msg.Fields {
			fldTypeFile, err := g.parsedFile(fld.Type.File)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to resolve message '%s' field '%s' type parsed file", dotedTypeName(msg.TypeName), fld)
			}
			fldCfg := msgCfg.Fields[fld.Name]
			resolver, withErr, err := g.TypeValueResolver(fldTypeFile, fld.Type, fldCfg.ContextKey)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get type value resolver")
			}
			goType, err := g.goTypeByParserType(fld.Type)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get go type by parser type")
			}
			if fld.Repeated {
				gt := goType
				goType = common.GoType{
					Kind:     reflect.Slice,
					ElemType: &gt,
				}
			}
			fields = append(fields, common.InputObjectResolverField{
				GraphQLInputFieldName: fld.Name,
				OutputFieldName:       camelCase(fld.Name),
				ValueResolver:         resolver,
				ResolverWithError:     withErr,
				GoType:                goType,
			})
		}
		for _, fld := range msg.MapFields {
			fldCfg := msgCfg.Fields[fld.Name]
			valueResolver, withErr, err := g.TypeValueResolver(file, fld.Type, fldCfg.ContextKey)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to get message '%s' map field '%s' value resolver", msg.Name, fld.Name)
			}
			fields = append(fields, common.InputObjectResolverField{
				GraphQLInputFieldName: fld.Name,
				OutputFieldName:       camelCase(fld.Name),
				ValueResolver:         valueResolver,
				ResolverWithError:     withErr,
			})
		}
		msgGoType, err := g.goTypeByParserType(msg.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve message go type")
		}
		res = append(res, common.InputObjectResolver{
			FunctionName: g.inputMessageResolverName(file, msg),
			OutputGoType: msgGoType,
			OneOfFields:  oneOffs,
			Fields:       fields,
		})

	}
	return res, nil
}
