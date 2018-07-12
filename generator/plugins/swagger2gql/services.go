package swagger2gql

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

func (p *Plugin) graphqlMethod(methodCfg MethodConfig, file *parsedFile, tag parser.Tag, method parser.Method) (graphql.Method, error) {
	name := method.OperationID
	if methodCfg.Alias != "" {
		name = methodCfg.Alias
	}
	var successResponse parser.MethodResponse
	var successResponseFound bool
	for _, resp := range method.Responses {
		if resp.StatusCode/100 == 2 {
			if successResponseFound {
				return graphql.Method{}, errors.New("method contains multiple success responses")
			}

			successResponse = resp
			successResponseFound = true
		}
	}
	responseType, err := p.TypeOutputTypeResolver(file, successResponse.ResultType, false)
	if err != nil {
		return graphql.Method{}, errors.Wrap(err, "can't get response output type resolver")
	}
	var args []graphql.MethodArgument
	for _, param := range method.Parameters {
		paramType, err := p.TypeInputTypeResolver(file, param.Type)
		if err != nil {
			return graphql.Method{}, errors.Wrapf(err, "failed to resolve parameter '%s' type resolver", param.Name)
		}
		args = append(args, graphql.MethodArgument{
			Name: pascalize(param.Name),
			Type: paramType,
		})
	}
	return graphql.Method{
		Name:              name,
		GraphQLOutputType: responseType,
		Arguments:         args,
		RequestResolver: func(arg string, ctx graphql.BodyContext) string {
			if ctx.TracerEnabled {
				return "Resolve" + pascalize(method.OperationID) + "Params(tr, tr.ContextWithSpan(ctx, span), " + arg + ")"
			}
			return "Resolve" + pascalize(method.OperationID) + "Params(ctx, " + arg + ")"
		},
		RequestResolverWithErr: true,
		ClientMethodCaller: func(arg string, ctx graphql.BodyContext) string {
			return pascalize(method.OperationID) + "(" + arg + ")"
		},
		RequestType: graphql.GoType{
			Kind: reflect.Ptr,
			ElemType: &graphql.GoType{
				Kind: reflect.Interface,
				Pkg:  file.Config.Tags[tag.Name].ClientGoPackage,
				Name: pascalize(method.OperationID) + "Params",
			},
		},
		PayloadErrorChecker:  nil,
		PayloadErrorAccessor: nil,
	}, nil
}
func (p *Plugin) tagQueriesMethods(tagCfg TagConfig, file *parsedFile, tag parser.Tag) ([]graphql.Method, error) {
	var res []graphql.Method
	for _, method := range tag.Methods {
		if method.HTTPMethod != "GET" {
			continue
		}
		var methodCfg MethodConfig
		if tagCfg.Methods[method.Path] != nil {
			methodCfg = tagCfg.Methods[method.Path][strings.ToLower(method.HTTPMethod)]
		}
		meth, err := p.graphqlMethod(methodCfg, file, tag, method)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve graphql method")
		}
		res = append(res, meth)
	}
	return res, nil
}
func (p *Plugin) tagMutationsMethods(tagCfg TagConfig, file *parsedFile, tag parser.Tag) ([]graphql.Method, error) {
	var res []graphql.Method
	for _, method := range tag.Methods {
		if method.HTTPMethod == "GET" {
			continue
		}
		var methodCfg MethodConfig
		if tagCfg.Methods[method.Path] != nil {
			methodCfg = tagCfg.Methods[method.Path][strings.ToLower(method.HTTPMethod)]
		}
		meth, err := p.graphqlMethod(methodCfg, file, tag, method)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve graphql method")
		}
		res = append(res, meth)
	}
	return res, nil
}
func (p *Plugin) fileServices(file *parsedFile) ([]graphql.Service, error) {
	var res []graphql.Service
	for _, tag := range file.File.Tags {
		tagCfg := file.Config.Tags[tag.Name]
		queriesMethods, err := p.tagQueriesMethods(tagCfg, file, tag)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get tag queries methods")
		}
		res = append(res, graphql.Service{
			Name:    pascalize(tag.Name),
			Methods: queriesMethods,
			CallInterface: graphql.GoType{
				Kind: reflect.Ptr,
				ElemType: &graphql.GoType{
					Kind: reflect.Interface,
					Pkg:  file.Config.Tags[tag.Name].ClientGoPackage,
					Name: "Client",
				},
			},
		})
		mutationsMethods, err := p.tagMutationsMethods(tagCfg, file, tag)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get tag queries methods")
		}
		res = append(res, graphql.Service{
			Name:    pascalize(tag.Name) + "Mutations",
			Methods: mutationsMethods,
			CallInterface: graphql.GoType{
				Kind: reflect.Ptr,
				ElemType: &graphql.GoType{
					Kind: reflect.Interface,
					Pkg:  file.Config.Tags[tag.Name].ClientGoPackage,
					Name: "Client",
				},
			},
		})
	}
	return res, nil
}
