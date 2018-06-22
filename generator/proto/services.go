package proto

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g Generator) serviceMethodArguments(file *parser.File, method *parser.Method) ([]common.MethodArgument, error) {
	var args []common.MethodArgument
	for _, field := range method.InputMessage.Fields {
		typResolver, err := g.TypeInputTypeResolver(file, field.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to prepare input type resolver")
		}
		args = append(args, common.MethodArgument{
			Name: field.Name,
			Type: typResolver,
		})
	}
	for _, field := range method.InputMessage.Fields {
		typResolver, err := g.TypeInputTypeResolver(file, field.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to prepare input type resolver")
		}
		args = append(args, common.MethodArgument{
			Name: field.Name,
			Type: typResolver,
		})
	}
	return args, nil
}
func (g Generator) serviceMethod(file *parser.File, method *parser.Method) (*common.Method, error) {
	outType, err := g.TypeOutputTypeResolver(file, method.OutputMessage.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get output type resolver for method: ", method.Name)
	}
	requestType, err := goTypeByParserType(method.InputMessage.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get request go type for method: ", method.Name)
	}
	responseType, err := goTypeByParserType(method.OutputMessage.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get request go type for method: ", method.Name)
	}
	args, err := g.serviceMethodArguments(file, method)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare service method arguments")
	}
	return &common.Method{
		Name:                        method.Name,
		GraphQLOutputType:           outType,
		RequestType:                 requestType,
		ResponseType:                responseType,
		CallMethod:                  camelCase(method.Name),
		RequestResolverFunctionName: g.inputMessageResolverName(method.InputMessage),
		Arguments:                   args,
	}, nil
}
func (g Generator) serviceQueryMethods(file *parser.File, service *parser.Service) ([]common.Method, error) {
	var res []common.Method
	for _, method := range service.Methods {
		if !strings.HasPrefix(strings.ToLower(method.Name), "get") {
			continue
		}
		met, err := g.serviceMethod(file, method)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare service method %s", method.Name)
		}

		res = append(res, *met)
	}
	return res, nil
}
func (g Generator) serviceMutationsMethods(file *parser.File, service *parser.Service) ([]common.Method, error) {
	var res []common.Method
	for _, method := range service.Methods {
		if strings.HasPrefix(strings.ToLower(method.Name), "get") {
			continue
		}
		met, err := g.serviceMethod(file, method)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare service method %s", method.Name)
		}

		res = append(res, *met)
	}
	return res, nil
}
func (g Generator) fileServices(file parsedFile) ([]common.Service, error) {
	var res []common.Service
	for _, service := range file.File.Services {
		queryMethods, err := g.serviceQueryMethods(file.File, service)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve service methods")
		}
		res = append(res, common.Service{
			Name: service.Name, // TODO: use aliase
			CallInterface: common.GoType{
				Kind: reflect.Interface,
				Pkg:  file.File.GoPackage,
				Name: service.Name + "Client",
			},
			Methods: queryMethods,
		})
		mutationsMethods, err := g.serviceMutationsMethods(file.File, service)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve service methods")
		}
		res = append(res, common.Service{
			Name: "Mutations" + service.Name, // TODO: use aliase
			CallInterface: common.GoType{
				Kind: reflect.Interface,
				Pkg:  file.File.GoPackage,
				Name: service.Name + "Client",
			},
			Methods: mutationsMethods,
		})
	}
	return res, nil
}
