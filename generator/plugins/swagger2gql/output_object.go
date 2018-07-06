package swagger2gql

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

func (g *Plugin) outputObjectGQLName(messageFile *parsedFile, obj parser.Object) string {
	return messageFile.Config.GetGQLMessagePrefix() + strings.Join(obj.Route, "__")
}
func (g *Plugin) outputObjectVariable(messageFile *parsedFile, obj parser.Object) string {
	return messageFile.Config.GetGQLMessagePrefix() + strings.Join(obj.Route, "")
}

func (g *Plugin) outputMessageTypeResolver(messageFile *parsedFile, obj parser.Object) (graphql.TypeResolver, error) {
	if len(obj.Properties) == 0 {
		return graphql.GqlNoDataTypeResolver, nil
	}
	return func(ctx graphql.BodyContext) string {
		return ctx.Importer.Prefix(messageFile.OutputPkg) + g.outputObjectVariable(messageFile, obj)
	}, nil
}

func (g *Plugin) outputMessageFields(file *parsedFile, obj parser.Object) ([]graphql.ObjectField, error) {
	var res []graphql.ObjectField
	for _, field := range obj.Properties {
		if _, ok := field.Type.(parser.Map); ok {
			continue
		}
		typeResolver, err := g.TypeOutputTypeResolver(file, field.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare message %s field %s output type resolver", obj.Name, field.Name)
		}
		res = append(res, graphql.ObjectField{
			Name:           field.Name,
			Type:           typeResolver,
			GoObjectGetter: camelCase(field.Name),
		})
	}
	return res, nil
}

func (g *Plugin) outputMessageMapFields(file *parsedFile, msg parser.Object) ([]graphql.ObjectField, error) {
	var res []graphql.ObjectField
	for _, property := range msg.Properties {
		if _, ok := property.Type.(parser.Map); !ok {
			continue
		}
		typeResolver, err := g.TypeOutputTypeResolver(file, property.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare message %s property %s output type resolver", msg.Name, property.Name)
		}
		res = append(res, graphql.ObjectField{
			Name:           property.Name,
			Type:           typeResolver,
			GoObjectGetter: camelCase(property.Name),
		})
	}
	return res, nil
}

func (g *Plugin) fileOutputMessages(file *parsedFile) ([]graphql.OutputObject, error) {
	var res []graphql.OutputObject
	handledObjects := map[string]struct{}{}
	var handleType func(typ parser.Type) error
	handleType = func(typ parser.Type) error {
		switch t := typ.(type) {
		case parser.Object:
			if _, handled := handledObjects[snakeCamelCaseSlice(t.Route)]; handled {
				return nil
			}
			for _, property := range t.Properties {
				if err := handleType(property.Type); err != nil {
					return errors.Wrapf(err, "failed to handle object property %s type", property.Name)
				}
			}
			goTyp, err := g.goTypeByParserType(file, t, false)
			if err != nil {
				return errors.Wrap(err, "failed to resolve object go type")
			}
			var fields []graphql.ObjectField
			for _, prop := range t.Properties {
				tr, err := g.TypeOutputTypeResolver(file, prop.Type)
				if err != nil {
					return errors.Wrap(err, "failed to resolve property output type resolver")
				}
				fields = append(fields, graphql.ObjectField{
					Name:           prop.Name,
					Type:           tr,
					GoObjectGetter: pascalize(prop.Name),
					NeedCast:       false,
				})
			}
			res = append(res, graphql.OutputObject{
				VariableName: g.outputObjectVariable(file, t),
				GraphQLName:  g.outputObjectGQLName(file, t),
				GoType:       goTyp,
				Fields:       fields,
			})
			handledObjects[snakeCamelCaseSlice(t.Route)] = struct{}{}
		case parser.Array:
			return handleType(t.ElemType)
		}
		return nil
	}
	for _, tag := range file.File.Tags {
		for _, method := range tag.Methods {
			for _, resp := range method.Responses {
				if err := handleType(resp.ResultType); err != nil {
					return nil, errors.Wrapf(err, "failed to handle %s method %d response", method.OperationID, resp.StatusCode)
				}

			}
		}
	}
	return res, nil
}
