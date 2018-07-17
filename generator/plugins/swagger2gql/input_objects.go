package swagger2gql

import (
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

func (p *Plugin) inputObjectGQLName(file *parsedFile, obj *parser.Object) string {
	return file.Config.GetGQLMessagePrefix() + pascalize(strings.Join(obj.Route, "__")) + "Input"
}
func (p *Plugin) inputObjectVariable(msgFile *parsedFile, obj *parser.Object) string {
	return msgFile.Config.GetGQLMessagePrefix() + pascalize(strings.Join(obj.Route, "")) + "Input"
}
func (p *Plugin) methodParamsInputObjectVariable(file *parsedFile, method parser.Method) string {
	return file.Config.GetGQLMessagePrefix() + pascalize(method.OperationID+"Params") + "Input"
}
func (p *Plugin) methodParamsInputObjectGQLName(file *parsedFile, method parser.Method) string {
	return file.Config.GetGQLMessagePrefix() + pascalize(method.OperationID+"Params") + "Input"
}

//
func (p *Plugin) inputObjectTypeResolver(msgFile *parsedFile, obj *parser.Object) graphql.TypeResolver {
	if len(obj.Properties) == 0 {
		return graphql.GqlNoDataTypeResolver
	}

	return func(ctx graphql.BodyContext) string {
		return ctx.Importer.Prefix(msgFile.OutputPkg) + p.inputObjectVariable(msgFile, obj)
	}
}

func (p *Plugin) methodParametersInputObject(methodCfg MethodConfig, file *parsedFile, tag string, method parser.Method) (graphql.InputObject, error) {
	var fields []graphql.ObjectField
	gqlName := p.methodParamsInputObjectGQLName(file, method)
	cfg, err := file.Config.ObjectConfig(gqlName)
	if err != nil {
		return graphql.InputObject{}, errors.Wrap(err, "failed to resolve object config")
	}
	for _, parameter := range method.Parameters {
		typResovler, err := p.TypeInputTypeResolver(file, parameter.Type)
		if err != nil {
			return graphql.InputObject{}, errors.Wrapf(err, "failed to resolve parameter %s type resolver", parameter.Name)
		}
		paramName := pascalize(parameter.Name)
		paramCfg, _ := cfg.Fields[paramName]

		if paramCfg.ContextKey != "" {
			continue
		}
		fields = append(fields, graphql.ObjectField{
			Name:     paramName,
			Type:     typResovler,
			Value:    graphql.IdentAccessValueResolver(camelCase(pascalize(parameter.Name))),
			NeedCast: false,
		})
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name > fields[j].Name
	})

	return graphql.InputObject{
		VariableName: p.methodParamsInputObjectVariable(file, method),
		GraphQLName:  gqlName,
		Fields:       fields,
	}, nil
}

func (p *Plugin) fileInputObjects(file *parsedFile) ([]graphql.InputObject, error) {
	var res []graphql.InputObject
	var handledObjects = map[parser.Type]struct{}{}
	var handleType func(typ parser.Type) error
	handleType = func(typ parser.Type) error {
		switch t := typ.(type) {
		case *parser.Object:
			if _, handled := handledObjects[typ]; handled {
				return nil
			}
			handledObjects[typ] = struct{}{}
			var fields []graphql.ObjectField
			for _, property := range t.Properties {
				if err := handleType(property.Type); err != nil {
					return err
				}
				typeResolver, err := p.TypeInputTypeResolver(file, property.Type)
				if err != nil {
					return errors.Wrap(err, "failed to get input type resolver")
				}
				if property.Required {
					typeResolver = graphql.GqlNonNullTypeResolver(typeResolver)
				}
				fields = append(fields, graphql.ObjectField{
					Name:     pascalize(property.Name),
					Type:     typeResolver,
					NeedCast: false,
				})
			}
			sort.Slice(fields, func(i, j int) bool {
				return fields[i].Name > fields[j].Name
			})
			res = append(res, graphql.InputObject{
				VariableName: p.inputObjectVariable(file, t),
				GraphQLName:  p.inputObjectGQLName(file, t),
				Fields:       fields,
			})

		case *parser.Array:
			return handleType(t.ElemType)
		}
		return nil
	}
	for _, tag := range file.File.Tags {
		tagCfg := file.Config.Tags[tag.Name]
		for _, method := range tag.Methods {
			methodCfg := tagCfg.Methods[method.Path][method.HTTPMethod]

			parametersObj, err := p.methodParametersInputObject(methodCfg, file, tag.Name, method)
			if err != nil {
				return nil, errors.Wrap(err, "failed to prepare method parameters input object")
			}
			res = append(res, parametersObj)
			for _, parameter := range method.Parameters {
				err := handleType(parameter.Type)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to handle method %s parameter %s", method.OperationID, parameter.Name)
				}
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].VariableName > res[j].VariableName
	})
	return res, nil
}
