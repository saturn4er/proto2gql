package proto

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

func (g Generator) serviceMethodArguments(cfg MethodConfig, file *parser.File, method *parser.Method) ([]common.MethodArgument, error) {
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
	for _, field := range method.InputMessage.MapFields {
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
func (g Generator) serviceMethod(cfg MethodConfig, file *parser.File, method *parser.Method) (*common.Method, error) {
	outType, err := g.TypeOutputTypeResolver(method.OutputMessage.Type)
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
	args, err := g.serviceMethodArguments(cfg, file, method)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare service method arguments")
	}
	name := method.Name
	if cfg.Alias != "" {
		name = cfg.Alias
	}
	return &common.Method{
		Name:                        name,
		GraphQLOutputType:           outType,
		RequestType:                 requestType,
		ResponseType:                responseType,
		CallMethod:                  camelCase(method.Name),
		RequestResolverFunctionName: g.inputMessageResolverName(method.InputMessage),
		Arguments:                   args,
	}, nil
}
func (g Generator) serviceQueryMethods(sc ServiceConfig, file *parser.File, service *parser.Service) ([]common.Method, error) {
	var res []common.Method
	for _, method := range service.Methods {
		mc := sc.Methods[method.Name]
		if !g.methodIsQuery(mc, method) {
			continue
		}
		met, err := g.serviceMethod(mc, file, method)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare service method %s", method.Name)
		}
		res = append(res, *met)
	}
	return res, nil
}
func (g Generator) methodIsQuery(cfg MethodConfig, method *parser.Method) bool {
	return !strings.HasPrefix(strings.ToLower(method.Name), "get")
}
func (g Generator) serviceMutationsMethods(sc ServiceConfig, file *parser.File, service *parser.Service) ([]common.Method, error) {
	var res []common.Method
	for _, method := range service.Methods {
		mc := sc.Methods[method.Name]
		if !g.methodIsMutation(mc, method) {
			continue
		}
		met, err := g.serviceMethod(mc, file, method)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to prepare service method %s", method.Name)
		}

		res = append(res, *met)
	}
	return res, nil
}
func (g Generator) methodIsMutation(cfg MethodConfig, method *parser.Method) bool {
	return strings.HasPrefix(strings.ToLower(method.Name), "get")
}
func (g Generator) fileServices(file *parser.File) ([]common.Service, error) {
	var res []common.Service
	cfg := g.fileConfig(file)
	for _, service := range file.Services {
		serviceName := service.Name
		sc := cfg.GetServices()[service.Name]
		if sc.Alias != "" {
			serviceName = sc.Alias
		}
		queryMethods, err := g.serviceQueryMethods(sc, file, service)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve service methods")
		}
		res = append(res, common.Service{
			Name: serviceName, // TODO: use aliase
			CallInterface: common.GoType{
				Kind: reflect.Interface,
				Pkg:  g.fileGRPCSourcesPackage(file),
				Name: serviceName + "Client",
			},
			Methods: queryMethods,
		})
		mutationsMethods, err := g.serviceMutationsMethods(sc, file, service)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve service methods")
		}
		res = append(res, common.Service{
			Name: "Mutations" + serviceName, // TODO: use aliase
			CallInterface: common.GoType{
				Kind: reflect.Interface,
				Pkg:  g.fileGRPCSourcesPackage(file),
				Name: serviceName + "Client",
			},
			Methods: mutationsMethods,
		})
	}
	return res, nil
}
