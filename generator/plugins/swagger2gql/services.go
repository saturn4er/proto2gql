package swagger2gql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

func (p *Plugin) graphqlMethod(file *parsedFile, tag parser.Tag, method parser.Method) graphql.Method {
	return graphql.Method{
		Name:              method.OperationID,
		GraphQLOutputType: func(ctx graphql.BodyContext) string {
			return p.methodParamsInputObjectVariable(file, method)
		},
		Arguments:         nil,
		RequestResolver: func(arg string, ctx graphql.BodyContext) string {
			if ctx.TracerEnabled {
				return "Resolve" +  pascalize(method.OperationID) + "Params(tr, tr.ContextWithSpan(ctx, span), " + arg + ")"
			}
			return "Resolve" +  pascalize(method.OperationID) + "Params(ctx, " + arg + ")"
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
	}
}
func (p *Plugin) tagQueriesMethods(file *parsedFile, tag parser.Tag) ([]graphql.Method, error) {
	var res []graphql.Method
	for _, method := range tag.Methods {
		if method.HTTPMethod != "GET" {
			continue
		}
		res = append(res, p.graphqlMethod(file, tag, method))
	}
	return res, nil
}
func (p *Plugin) tagMutationsMethods(file *parsedFile, tag parser.Tag) ([]graphql.Method, error) {
	var res []graphql.Method
	for _, method := range tag.Methods {
		if method.HTTPMethod == "GET" {
			continue
		}
		res = append(res, p.graphqlMethod(file, tag, method))
	}
	return res, nil
}
func (p *Plugin) fileServices(file *parsedFile) ([]graphql.Service, error) {
	var res []graphql.Service
	for _, tag := range file.File.Tags {
		queriesMethods, err := p.tagQueriesMethods(file, tag)
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
		mutationsMethods, err := p.tagMutationsMethods(file, tag)
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
