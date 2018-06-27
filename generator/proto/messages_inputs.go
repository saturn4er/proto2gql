package proto

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) inputMessageGraphQLName(file *parsedFile, message *parser.Message) string {
	return file.Config.GetGQLMessagePrefix() + strings.Join(message.TypeName, "__") + "Input"
}
func (g *Generator) inputMessageVariable(msgFile *parsedFile, message *parser.Message) string {
	return msgFile.Config.GetGQLMessagePrefix() + strings.Join(message.TypeName, "") + "Input"
}

func (g *Generator) inputMessageTypeResolver(msgFile *parsedFile, message *parser.Message) (common.TypeResolver, error) {
	if !message.HaveFields() {
		return common.GqlNoDataTypeResolver, nil
	}

	return func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(msgFile.OutputPkg) + g.inputMessageVariable(msgFile, message)
	}, nil
}

func (g *Generator) inputMessageFieldTypeResolver(file *parsedFile, field *parser.Field) (common.TypeResolver, error) {
	resolver, err := g.TypeOutputTypeResolver(file, field.Type)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get input type resolver")
	}
	if field.Repeated {
		resolver = common.GqlListTypeResolver(common.GqlNonNullTypeResolver(resolver))
	}
	return resolver, nil
}

func (g *Generator) outputObjectMapFieldTypeResolver(mapFile *parsedFile, mp *parser.Map) (common.TypeResolver, error) {
	res := func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(mapFile.OutputPkg) + g.outputMapVariable(mapFile, mp)
	}
	return common.GqlListTypeResolver(common.GqlNonNullTypeResolver(res)), nil
}
func (g *Generator) inputObjectMapFieldTypeResolver(mapFile *parsedFile, mp *parser.Map) (common.TypeResolver, error) {
	res := func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(mapFile.OutputPkg) + g.inputMapVariable(mapFile, mp)
	}
	return common.GqlListTypeResolver(common.GqlNonNullTypeResolver(res)), nil
}

func (g *Generator) fileInputObjects(file *parsedFile) ([]common.InputObject, error) {
	var res []common.InputObject
	for _, msg := range file.File.Messages {
		if !msg.HaveFields() {
			continue
		}
		var fields []common.ObjectField
		for _, field := range msg.Fields {
			fieldTypeFile, err := g.parsedFile(field.Type.File)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value type file")
			}
			typ, err := g.inputMessageFieldTypeResolver(fieldTypeFile, field)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve field type")
			}
			fields = append(fields, common.ObjectField{
				Name: field.Name,
				Type: typ,
			})
		}
		for _, field := range msg.MapFields {
			typ, err := g.inputObjectMapFieldTypeResolver(file, field.Map)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve field type")
			}
			fields = append(fields, common.ObjectField{
				Name: field.Name,
				Type: typ,
			})
		}
		for _, oneOf := range msg.OneOffs {
			for _, fld := range oneOf.Fields {
				fieldTypeFile, err := g.parsedFile(fld.Type.File)
				if err != nil {
					return nil, errors.Wrap(err, "failed to resolve value type file")
				}
				typ, err := g.inputMessageFieldTypeResolver(fieldTypeFile, fld)
				if err != nil {
					return nil, errors.Wrap(err, "failed to resolve field type")
				}
				fields = append(fields, common.ObjectField{
					Name: fld.Name,
					Type: typ,
				})
			}
		}
		// TODO: oneof fields
		res = append(res, common.InputObject{
			VariableName: g.inputMessageVariable(file, msg),
			GraphQLName:  g.inputMessageGraphQLName(file, msg),
			Fields:       fields,
		})
	}
	return res, nil
}
