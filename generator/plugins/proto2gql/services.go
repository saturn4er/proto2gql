package proto2gql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/proto2gql/parser"
)

// func (g Proto2GraphQL) SchemaServices() []schema.Service {
// 	var res []schema.Service
// 	for _, file := range g.ParsedFiles {
// 		for _, service := range file.TypesFile.Services {
// 			sc := file.Config.GetServices()[service.Name]
// 			clientGoType := schema.GoType{
// 				Kind: reflect.Interface,
// 				Pkg:  file.GRPCSourcesPkg,
// 				Name: service.Name + "Client",
// 			}
// 			queryService := schema.Service{
// 				Name:          g.serviceQueryName(sc, service),
// 				Pkg:           file.OutputPkg,
// 				ClientGoType:  clientGoType,
// 				TracerEnabled: g.GenerateTracers,
// 			}
// 			mutationService := schema.Service{
// 				Name:          g.serviceMutationName(sc, service),
// 				Pkg:           file.OutputPkg,
// 				ClientGoType:  clientGoType,
// 				TracerEnabled: g.GenerateTracers,
// 			}
// 			for _, method := range service.Methods {
// 				methodCfg := sc.Methods[method.Name]
// 				if g.methodIsQuery(methodCfg, method) {
// 					queryService.Fields = append(queryService.Fields, g.methodName(methodCfg, method))
// 				}
// 				if g.methodIsMutation(methodCfg, method) {
// 					mutationService.Fields = append(mutationService.Fields, g.methodName(methodCfg, method))
// 				}
// 			}
//
// 			res = append(res, queryService, mutationService)
// 		}
// 	}
// 	return res
// }
func (g Proto2GraphQL) serviceMethodArguments(cfg MethodConfig, method *parser.Method) ([]graphql.MethodArgument, error) {
	var args []graphql.MethodArgument
	for _, field := range method.InputMessage.Fields {
		typeFile, err := g.parsedFile(field.Type.File)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve field '%s' file", field.Name)
		}
		typResolver, err := g.TypeInputTypeResolver(typeFile, field.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to prepare input type resolver")
		}
		args = append(args, graphql.MethodArgument{
			Name: field.Name,
			Type: typResolver,
		})
	}
	for _, field := range method.InputMessage.MapFields {
		typeFile, err := g.parsedFile(field.Type.File)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve field '%s' file", field.Name)
		}
		typResolver, err := g.TypeInputTypeResolver(typeFile, field.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to prepare input type resolver")
		}
		args = append(args, graphql.MethodArgument{
			Name: field.Name,
			Type: typResolver,
		})
	}
	return args, nil
}
func (g Proto2GraphQL) messagePayloadErrorParams(message *parser.Message) (checker graphql.PayloadErrorChecker, accessor graphql.PayloadErrorAccessor, err error) {
	outMsgCfg, err := g.fileConfig(message.File).MessageConfig(dotedTypeName(message.TypeName))
	if err != nil {
		err = errors.Wrap(err, "failed to resolve output message config")
		return
	}
	if outMsgCfg.ErrorField == "" {
		return
	}
	errorAccessor := func(arg string) string {
		return arg + ".Get" + camelCase(outMsgCfg.ErrorField) + "()"
	}
	errorCheckerByType := func(repeated bool, p *parser.Type) graphql.PayloadErrorChecker {
		if repeated || p.IsMap() {
			return func(arg string) string {
				return "len(" + arg + ".Get" + camelCase(outMsgCfg.ErrorField) + "())>0"
			}
		}
		if p.IsScalar() || p.IsEnum() {
			fmt.Println("Warning: scalars and enums is not supported as payload error fields")
			return nil
		}
		if p.IsMessage() {
			return func(arg string) string {
				return arg + ".Get" + camelCase(outMsgCfg.ErrorField) + "() != nil"
			}
		}
		return nil
	}
	for _, fld := range message.Fields {
		if fld.Name == outMsgCfg.ErrorField {
			errorChecker := errorCheckerByType(fld.Repeated, fld.Type)
			if errorChecker == nil {
				return nil, nil, nil
			}
			return errorChecker, errorAccessor, nil
		}
	}
	for _, fld := range message.MapFields {
		if fld.Name == outMsgCfg.ErrorField {
			errorChecker := errorCheckerByType(false, fld.Type)
			if errorChecker == nil {
				return nil, nil, nil
			}
			return errorChecker, errorAccessor, nil
		}
	}
	for _, of := range message.OneOffs {
		for _, fld := range of.Fields {
			if fld.Name == outMsgCfg.ErrorField {
				errorChecker := errorCheckerByType(false, fld.Type)
				if errorChecker == nil {
					return nil, nil, nil
				}
				return errorChecker, errorAccessor, nil
			}
		}
	}
	return nil, nil, nil
}
func (g Proto2GraphQL) methodName(cfg MethodConfig, method *parser.Method) string {
	if cfg.Alias != "" {
		return cfg.Alias
	}
	return method.Name
}
func (g Proto2GraphQL) serviceMethod(cfg MethodConfig, file *parsedFile, method *parser.Method) (*graphql.Method, error) {
	outputMsgTypeFile, err := g.parsedFile(method.OutputMessage.File)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file type file")
	}
	outType, err := g.TypeOutputTypeResolver(outputMsgTypeFile, method.OutputMessage.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get output type resolver for method: ", method.Name)
	}
	requestType, err := g.goTypeByParserType(method.InputMessage.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get request go type for method: ", method.Name)
	}
	args, err := g.serviceMethodArguments(cfg, method)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare service method arguments")
	}
	payloadErrChecker, payloadErrAccessor, err := g.messagePayloadErrorParams(method.OutputMessage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve message payload error params")
	}
	inputMessageFile, err := g.parsedFile(method.InputMessage.File)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve message '%s' parsed file", dotedTypeName(method.InputMessage.TypeName))
	}
	valueResolver, valueResolverWithErr, err := g.TypeValueResolver(inputMessageFile, method.InputMessage.Type, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve message value resolver")
	}
	return &graphql.Method{
		Name:                   g.methodName(cfg, method),
		GraphQLOutputType:      outType,
		RequestType:            requestType,
		CallMethod:             camelCase(method.Name),
		RequestResolver:        valueResolver,
		RequestResolverWithErr: valueResolverWithErr,
		Arguments:              args,
		PayloadErrorChecker:    payloadErrChecker,
		PayloadErrorAccessor:   payloadErrAccessor,
	}, nil
}
func (g Proto2GraphQL) serviceQueryMethods(sc ServiceConfig, file *parsedFile, service *parser.Service) ([]graphql.Method, error) {
	var res []graphql.Method
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
func (g Proto2GraphQL) methodIsQuery(cfg MethodConfig, method *parser.Method) bool {
	return !strings.HasPrefix(strings.ToLower(method.Name), "get")
}
func (g Proto2GraphQL) serviceMutationsMethods(cfg ServiceConfig, file *parsedFile, service *parser.Service) ([]graphql.Method, error) {
	var res []graphql.Method
	for _, method := range service.Methods {
		mc := cfg.Methods[method.Name]
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
func (g Proto2GraphQL) methodIsMutation(cfg MethodConfig, method *parser.Method) bool {
	return strings.HasPrefix(strings.ToLower(method.Name), "get")
}
func (g Proto2GraphQL) serviceQueryName(sc ServiceConfig, service *parser.Service) string {
	if sc.Alias != "" {
		return sc.Alias
	}
	return service.Name
}
func (g Proto2GraphQL) serviceMutationName(sc ServiceConfig, service *parser.Service) string {
	if sc.Alias != "" {
		return sc.Alias + "Mutations"
	}
	return service.Name + "Mutations"
}
func (g Proto2GraphQL) fileServices(file *parsedFile) ([]graphql.Service, error) {
	var res []graphql.Service
	for _, service := range file.File.Services {
		sc := file.Config.GetServices()[service.Name]
		queryMethods, err := g.serviceQueryMethods(sc, file, service)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve service methods")
		}
		res = append(res, graphql.Service{
			Name: g.serviceQueryName(sc, service),
			CallInterface: graphql.GoType{
				Kind: reflect.Interface,
				Pkg:  file.GRPCSourcesPkg,
				Name: service.Name + "Client",
			},
			Methods: queryMethods,
		})
		mutationsMethods, err := g.serviceMutationsMethods(sc, file, service)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve service methods")
		}
		res = append(res, graphql.Service{
			Name: g.serviceMutationName(sc, service),
			CallInterface: graphql.GoType{
				Kind: reflect.Interface,
				Pkg:  file.GRPCSourcesPkg,
				Name: service.Name + "Client",
			},
			Methods: mutationsMethods,
		})
	}
	return res, nil
}
