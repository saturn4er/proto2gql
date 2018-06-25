package proto

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) inputMessageResolverName(message *parser.Message) string {
	return "Resolve" + g.inputMessageGraphQLName(message)
}

func (g *Generator) oneOfValueAssigningWrapper(msg *parser.Message, field *parser.Field) common.AssigningWrapper {
	return func(arg string, ctx common.BodyContext) string {
		pkg := g.fileGRPCSourcesPackage(msg.Type.File)
		return "&" + ctx.Importer.Prefix(pkg) + camelCaseSlice(msg.TypeName) + "_" + camelCase(field.Name) + "{" + arg + "}"
	}
}

func (g *Generator) fileInputMessagesResolvers(file *parser.File) ([]common.InputObjectResolver, error) {
	var res []common.InputObjectResolver
	fileCfg := g.fileConfig(file)
	for _, msg := range file.Messages {
		msgCfg, err := fileCfg.MessageConfig(strings.Join(msg.TypeName, "."))
		if err != nil {

		}
		var oneOffs []common.InputObjectResolverOneOf
		for _, oneOf := range msg.OneOffs {
			var fields []common.InputObjectResolverOneOfField
			for _, fld := range oneOf.Fields {
				fldCfg := msgCfg.Fields[fld.Name]
				resolver, withErr, err := g.TypeValueResolver(fld.Type, fldCfg.ContextKey)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get type value resolver")
				}
				fields = append(fields, common.InputObjectResolverOneOfField{
					GraphQLInputFieldName: fld.Name,
					ValueResolver:         resolver,
					ResolverWithError:     withErr,
					AssigningWrapper:      g.oneOfValueAssigningWrapper(msg, fld),
				})
			}
			oneOffs = append(oneOffs, common.InputObjectResolverOneOf{
				OutputFieldName: camelCase(oneOf.Name),
				Fields:          fields,
			})
		}
		var fields []common.InputObjectResolverField
		for _, fld := range msg.Fields {
			fldCfg := msgCfg.Fields[fld.Name]
			resolver, withErr, err := g.TypeValueResolver(fld.Type, fldCfg.ContextKey)
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
			valueResolver, withErr, err := g.TypeValueResolver(fld.Type, fldCfg.ContextKey)
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
			FunctionName: g.inputMessageResolverName(msg),
			OutputGoType: msgGoType,
			OneOfFields:  oneOffs,
			Fields:       fields,
		})

	}
	return res, nil
}
