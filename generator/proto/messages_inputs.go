package proto

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)


func (g *Generator) inputMessageGraphQLName(message *parser.Message) string {
	return g.fileConfig(message.File).GetGQLMessagePrefix() + strings.Join(message.TypeName, "__") + "Input"
}
func (g *Generator) inputMessageVariable(message *parser.Message) string {
	return g.fileConfig(message.File).GetGQLMessagePrefix() + strings.Join(message.TypeName, "") + "Input"
}

func (g *Generator) inputMessageTypeResolver(currentFile *parser.File, message *parser.Message) (common.TypeResolver, error) {
	if !message.HaveFields() {
		return common.GqlNoDataTypeResolver, nil
	}
	_, pkg, err := g.fileOutputPackage(message.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file output package")
	}
	return func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(pkg) + g.inputMessageVariable(message)
	}, nil
}

func (g *Generator) inputMessageFieldTypeResolver(currentFile *parser.File, field *parser.Field) (common.TypeResolver, error) {
	resolver, err := g.TypeOutputTypeResolver(currentFile, field.Type)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get input type resolver")
	}
	if field.Repeated {
		resolver = common.GqlListTypeResolver(common.GqlNonNullTypeResolver(resolver))
	}
	return resolver, nil
}

func (g *Generator) inputObjectMapFieldTypeResolver(message *parser.Message, field *parser.MapField) (common.TypeResolver, error) {
	if !field.Type.IsMap() {
		return nil, errors.New("map field is not of type 'Map'")
	}
	_, pkg, err := g.fileOutputPackage(field.Type.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file output package")
	}
	res := func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(pkg) + g.inputMapVariable(field.Map)
	}
	return common.GqlListTypeResolver(common.GqlNonNullTypeResolver(res)), nil
}

func (g *Generator) fileInputObjects(file parsedFile) ([]common.InputObject, error) {
	var res []common.InputObject
	for _, msg := range file.File.Messages {
		if !msg.HaveFields() {
			continue
		}
		var fields []common.ObjectField
		for _, field := range msg.Fields {
			typ, err := g.inputMessageFieldTypeResolver(msg.File, field)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve field type")
			}
			fields = append(fields, common.ObjectField{
				Name: field.Name,
				Type: typ,
			})
		}
		for _, field := range msg.MapFields {
			typ, err := g.inputObjectMapFieldTypeResolver(msg, field)
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
				typ, err := g.inputMessageFieldTypeResolver(msg.File, fld)
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
			VariableName: g.inputMessageVariable(msg),
			GraphQLName:  g.inputMessageGraphQLName(msg),
			Fields:       fields,
		})
	}
	return res, nil
}
