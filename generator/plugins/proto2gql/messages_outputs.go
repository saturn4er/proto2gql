package proto2gql

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/proto2gql/parser"
)

func (g *Proto2GraphQL) outputMessageGraphQLName(messageFile *parsedFile, message *parser.Message) string {
	return messageFile.Config.GetGQLMessagePrefix() + strings.Join(message.TypeName, "__")
}
func (g *Proto2GraphQL) outputMessageVariable(messageFile *parsedFile, message *parser.Message) string {
	return messageFile.Config.GetGQLMessagePrefix() + strings.Join(message.TypeName, "")
}

func (g *Proto2GraphQL) outputMessageTypeResolver(messageFile *parsedFile, message *parser.Message) (graphql.TypeResolver, error) {
	if !message.HaveFields() {
		return graphql.GqlNoDataTypeResolver, nil
	}
	return func(ctx graphql.BodyContext) string {
		return ctx.Importer.Prefix(messageFile.OutputPkg) + g.outputMessageVariable(messageFile, message)
	}, nil
}

func (g *Proto2GraphQL) outputMessageFields(file *parsedFile, msg *parser.Message) ([]graphql.ObjectField, error) {
	var res []graphql.ObjectField
	for _, field := range msg.Fields {
		fieldTypeFile, err := g.parsedFile(field.Type.File)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve file type file")
		}
		typeResolver, err := g.TypeOutputTypeResolver(fieldTypeFile, field.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare message %s field %s output type resolver", msg.Name, field.Name)
		}
		res = append(res, graphql.ObjectField{
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
			res = append(res, graphql.ObjectField{
				Name:           field.Name,
				Type:           typeResolver,
				GoObjectGetter: "Get" + camelCase(field.Name) + "()",
			})
		}
	}
	return res, nil
}

func (g *Proto2GraphQL) outputMessageMapFields(file *parsedFile, msg *parser.Message) ([]graphql.ObjectField, error) {
	var res []graphql.ObjectField
	for _, field := range msg.MapFields {
		typeResolver, err := g.TypeOutputTypeResolver(file, field.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare message %s field %s output type resolver", msg.Name, field.Name)
		}
		res = append(res, graphql.ObjectField{
			Name:           field.Name,
			Type:           typeResolver,
			GoObjectGetter: camelCase(field.Name),
		})
	}
	return res, nil
}

func (g *Proto2GraphQL) fileOutputMessages(file *parsedFile) ([]graphql.OutputObject, error) {
	var res []graphql.OutputObject
	for _, msg := range file.File.Messages {
		fields, err := g.outputMessageFields(file, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve message %s fields", msg.Name)
		}
		mapFields, err := g.outputMessageMapFields(file, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve message %s map fields", msg.Name)
		}
		res = append(res, graphql.OutputObject{
			VariableName: g.outputMessageVariable(file, msg),
			GraphQLName:  g.outputMessageGraphQLName(file, msg),
			Fields:       fields,
			MapFields:    mapFields,
			GoType: graphql.GoType{
				Kind: reflect.Struct,
				Name: snakeCamelCaseSlice(msg.TypeName),
				Pkg:  file.GRPCSourcesPkg,
			},
		})
	}
	return res, nil
}
