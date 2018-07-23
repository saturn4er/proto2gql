package swagger2gql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

func (p *Plugin) fileInputMapResolvers(file *parsedFile) ([]graphql.MapInputObjectResolver, error) {
	var res []graphql.MapInputObjectResolver
	handledObjects := map[parser.Type]struct{}{}
	var handleType func(typ parser.Type) error
	handleType = func(typ parser.Type) error {
		switch t := typ.(type) {
		case *parser.Map:
			valueGoType, err := p.goTypeByParserType(file, t.ElemType, false)
			if err != nil {
				return errors.Wrap(err, "failed to resolve map value go type")
			}
			valueResolver, valueWithErr, _, err := p.TypeValueResolver(file, t.ElemType, false, "")
			res = append(res, graphql.MapInputObjectResolver{
				FunctionName: "Resolve" + p.mapInputObjectVariable(file, t),
				KeyGoType: graphql.GoType{
					Kind: reflect.String,
				},
				ValueGoType: valueGoType,
				KeyResolver: func(arg string, ctx graphql.BodyContext) string {
					return arg + ".(string)"
				},
				KeyResolverWithError:   false,
				ValueResolver:          valueResolver,
				ValueResolverWithError: valueWithErr,
			})
		case *parser.Object:
			if _, handled := handledObjects[t]; handled {
				return nil
			}
			handledObjects[t] = struct{}{}
			for _, property := range t.Properties {
				if err := handleType(property.Type); err != nil {
					return errors.Wrapf(err, "failed to handle object property %s type", property.Name)
				}
			}
		case *parser.Array:
			return handleType(t.ElemType)
		}
		return nil
	}
	for _, tag := range file.File.Tags {
		for _, method := range tag.Methods {
			for _, param := range method.Parameters {
				if err := handleType(param.Type); err != nil {
					return nil, errors.Wrapf(err, "failed to handle %s method %s parameter", method.OperationID, param.Name)
				}

			}
		}
	}
	return res, nil
}
