package proto

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) outputMessageGraphQLName(messageFile *parsedFile, message *parser.Message) string {
	return messageFile.Config.GetGQLMessagePrefix() + strings.Join(message.TypeName, "__")
}
func (g *Generator) outputMessageVariable(messageFile *parsedFile, message *parser.Message) string {
	return messageFile.Config.GetGQLMessagePrefix() + strings.Join(message.TypeName, "")
}

func (g *Generator) outputMessageTypeResolver(messageFile *parsedFile, message *parser.Message) (common.TypeResolver, error) {
	if !message.HaveFields() {
		return common.GqlNoDataTypeResolver, nil
	}
	return func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(messageFile.OutputPkg) + g.outputMessageVariable(messageFile, message)
	}, nil
}

func (g *Generator) outputMessageFields(file *parsedFile, msg *parser.Message) ([]common.ObjectField, error) {
	var res []common.ObjectField
	for _, field := range msg.Fields {
		fieldTypeFile, err := g.parsedFile(field.Type.File)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve file type file")
		}
		typeResolver, err := g.TypeOutputTypeResolver(fieldTypeFile, field.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare message %s field %s output type resolver", msg.Name, field.Name)
		}
		res = append(res, common.ObjectField{
			Name:           field.Name,
			Type:           typeResolver,
			GoObjectGetter: camelCase(field.Name),
		})
	}
	for _, of := range msg.OneOffs {
		for _, field := range of.Fields {
			fieldTypeFile, err := g.parsedFile(field.Type.File)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve file type file")
			}
			typeResolver, err := g.TypeOutputTypeResolver(fieldTypeFile, field.Type)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to prepare message %s field %s output type resolver", msg.Name, field.Name)
			}
			res = append(res, common.ObjectField{
				Name:           field.Name,
				Type:           typeResolver,
				GoObjectGetter: "Get" + camelCase(field.Name) + "()",
			})
		}
	}
	return res, nil
}

func (g *Generator) outputMessageMapFields(file *parsedFile, msg *parser.Message) ([]common.ObjectField, error) {
	var res []common.ObjectField
	for _, field := range msg.MapFields {
		typeResolver, err := g.TypeOutputTypeResolver(file, field.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare message %s field %s output type resolver", msg.Name, field.Name)
		}
		res = append(res, common.ObjectField{
			Name:           field.Name,
			Type:           typeResolver,
			GoObjectGetter: camelCase(field.Name),
		})
	}
	return res, nil
}

func (g *Generator) fileOutputMessages(file *parsedFile) ([]common.OutputObject, error) {
	var res []common.OutputObject
	for _, msg := range file.File.Messages {
		fields, err := g.outputMessageFields(file, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve message %s fields", msg.Name)
		}
		mapFields, err := g.outputMessageMapFields(file, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve message %s map fields", msg.Name)
		}
		res = append(res, common.OutputObject{
			VariableName: g.outputMessageVariable(file, msg),
			GraphQLName:  g.outputMessageGraphQLName(file, msg),
			Fields:       fields,
			MapFields:    mapFields,
			GoType: common.GoType{
				Kind: reflect.Struct,
				Name: snakeCamelCaseSlice(msg.TypeName),
				Pkg:  file.GRPCSourcesPkg,
			},
		})
	}
	return res, nil
}
