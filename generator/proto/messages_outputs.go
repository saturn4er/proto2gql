package proto

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g *Generator) outputMessageGraphQLName(message *parser.Message) string {
	return g.fileConfig(message.File).GetGQLMessagePrefix() + strings.Join(message.TypeName, "__")
}
func (g *Generator) outputMessageVariable(message *parser.Message) string {
	return g.fileConfig(message.File).GetGQLMessagePrefix() + strings.Join(message.TypeName, "")
}

func (g *Generator) outputMessageTypeResolver(currentFile *parser.File, message *parser.Message) (common.TypeResolver, error) {
	if !message.HaveFields() {
		return common.GqlNoDataTypeResolver, nil
	}
	_, pkg, err := g.fileOutputPackage(message.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file output package")
	}
	return func(ctx common.BodyContext) string {
		return ctx.Importer.Prefix(pkg) + g.outputMessageVariable(message)
	}, nil
}

func (g *Generator) outputMessageFields(file parsedFile, msg *parser.Message) ([]common.ObjectField, error) {
	var res []common.ObjectField
	for _, field := range msg.Fields {
		typeResolver, err := g.TypeOutputTypeResolver(file.File, field.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare message %s field %s output type resolver", msg.Name, field.Name)
		}
		res = append(res, common.ObjectField{
			Name:          field.Name,
			Type:          typeResolver,
			GoObjectField: camelCase(field.Name),
		})
	}
	return res, nil
}
func (g *Generator) fileOutputMessages(file parsedFile) ([]common.OutputObject, error) {
	var res []common.OutputObject
	for _, msg := range file.File.Messages {
		fields, err := g.outputMessageFields(file, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve message %s fields", msg.Name)
		}
		res = append(res, common.OutputObject{
			VariableName: g.outputMessageVariable(msg),
			GraphQLName:  g.outputMessageGraphQLName(msg),
			Fields:       fields,
			GoType: common.GoType{
				Kind: reflect.Struct,
				Name: snakeCamelCaseSlice(msg.TypeName),
				Pkg:  g.fileGRPCSourcesPackage(file.File),
			},
		})
	}
	return res, nil
}
