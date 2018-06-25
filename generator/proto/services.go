package proto

import (
	"fmt"
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
func (g Generator) messageContainsField(message *parser.Message, field string) bool {

	return false
}
func (g Generator) messagePayloadErrorParams(message *parser.Message) (checker common.PayloadErrorChecker, accessor common.PayloadErrorAccessor, err error) {
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
	errorCheckerByType := func(repeated bool, p *parser.Type) common.PayloadErrorChecker {
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

func (g Generator) serviceMethod(cfg MethodConfig, file *parser.File, method *parser.Method) (*common.Method, error) {
	outType, err := g.TypeOutputTypeResolver(method.OutputMessage.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get output type resolver for method: ", method.Name)
	}
	requestType, err := g.goTypeByParserType(method.InputMessage.Type)
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
	payloadErrChecker, payloadErrAccessor, err := g.messagePayloadErrorParams(method.OutputMessage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve message payload error params")
	}
	return &common.Method{
		Name:                        name,
		GraphQLOutputType:           outType,
		RequestType:                 requestType,
		CallMethod:                  camelCase(method.Name),
		RequestResolverFunctionName: g.inputMessageResolverName(method.InputMessage),
		Arguments:                   args,
		PayloadErrorChecker:         payloadErrChecker,
		PayloadErrorAccessor:        payloadErrAccessor,
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
func (g Generator) serviceMutationsMethods(cfg ServiceConfig, file *parser.File, service *parser.Service) ([]common.Method, error) {
	var res []common.Method
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
