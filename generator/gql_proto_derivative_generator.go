package generator

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/parser"
	"golang.org/x/tools/imports"
)

const (
	scalarsPkgPath      = "github.com/saturn4er/proto2gql/api/scalars"
	interceptorsPkgPath = "github.com/saturn4er/proto2gql/api/interceptors"
	graphqlPkgPath      = "github.com/graphql-go/graphql"
	opentracingPkgPath  = "github.com/opentracing/opentracing-go"
	tracerPkg           = "github.com/saturn4er/proto2gql/api/tracer"
)

var customScalarsAnalogues = map[string]string{
	"double":   "GraphQLFloat64Scalar",
	"float":    "GraphQLFloat32Scalar",
	"int64":    "GraphQLInt64Scalar",
	"uint64":   "GraphQLUInt64Scalar",
	"int32":    "GraphQLInt32Scalar",
	"uint32":   "GraphQLUInt32Scalar",
	"fixed64":  "GraphQLUInt64Scalar",
	"fixed32":  "GraphQLUInt32Scalar",
	"bytes":    "GraphQLBytesScalar",
	"sfixed32": "GraphQLInt32Scalar",
	"sfixed64": "GraphQLInt64Scalar",
	"sint32":   "GraphQLInt32Scalar",
	"sint64":   "GraphQLInt64Scalar",
}
var graphqlScalarsAnalogues = map[string]string{
	"bool":   "Boolean",
	"string": "String",
}

type gqlProtoDerivativeFileGenerator struct {
	file           *gqlProtoDerivativeFile
	imports        importer
	generatedFiles []*gqlProtoDerivativeFile
}
type ErrorField struct {
	Name     string
	Repeated bool
	Type     *parser.Type
}

func (g *gqlProtoDerivativeFileGenerator) gqlEnumVarName(i *parser.Enum) string {
	return g.file.GQLEnumsPrefix + i.Name
}
func (g *gqlProtoDerivativeFileGenerator) scalarGoType(typ string) (string, error) {
	switch typ {
	case "double":
		return "float64", nil
	case "float":
		return "float32", nil
	case "int64", "sfixed64", "sint64":
		return "int64", nil
	case "int32", "sfixed32", "sint32":
		return "int32", nil
	case "uint64", "fixed64":
		return "uint64", nil
	case "uint32", "fixed32":
		return "uint32", nil
	case "bool":
		return "bool", nil
	case "string":
		return "string", nil
	case "bytes":
		return "[]byte", nil
	}
	return "", errors.New("not found")
}

func (g *gqlProtoDerivativeFileGenerator) scalarGQLType(imports *importer) func(typ string) (string, error) {
	return func(typ string) (string, error) {
		if v, ok := customScalarsAnalogues[typ]; ok {
			return imports.New(scalarsPkgPath) + "." + v, nil
		}
		if v, ok := graphqlScalarsAnalogues[typ]; ok {
			return imports.New(graphqlPkgPath) + "." + v, nil
		}
		return "", errors.Errorf("%s is not scalar", typ)
	}
}
func (g *gqlProtoDerivativeFileGenerator) gqlInputTypeName(imports *importer) (resolver func(t *parser.Type) string) {
	return func(t *parser.Type) string {
		if t.Message != nil && !t.Message.HaveFields() {
			return imports.New(scalarsPkgPath) + ".NoDataScalar"
		}
		if t.IsScalar() {
			res, err := g.scalarGQLType(imports)(t.Scalar)
			if err != nil {
				panic(err.Error())
			}
			return res
		}
		if t.File != g.file.ProtoFile {
			gf, err := g.findGeneratedFile(t.File)
			if err != nil {
				panic("Can't find generated import" + err.Error())
			}
			if gf.OutGoPkg != g.file.OutGoPkg {
				return imports.New(gf.OutGoPkg) + "." + gf.Generator.gqlInputTypeName(imports)(t)
			}
		}
		switch {
		case t.IsMessage():
			return g.file.GQLMessagePrefix + camelCaseSlice(t.Message.TypeName) + "Input"
		case t.IsEnum():
			return g.file.GQLEnumsPrefix + camelCaseSlice(t.Enum.TypeName)
		case t.IsMap():
			return g.file.GQLMessagePrefix + camelCaseSlice(t.Map.Message.TypeName) + camelCase(t.Map.Field.Name) + "MapInput"
		}
		panic(t.String() + " is not handled in gqlInputTypeName")
	}
}
func (g *gqlProtoDerivativeFileGenerator) gqlOutputTypeName(imports *importer) (res func(t *parser.Type) string) {
	return func(t *parser.Type) string {
		if t.Message != nil && !g.messageHaveFieldsExceptError(t.Message) {
			return imports.New(scalarsPkgPath) + ".NoDataScalar"
		}
		if t.IsScalar() {
			res, err := g.scalarGQLType(imports)(t.Scalar)
			if err != nil {
				panic(err.Error())
			}
			return res
		}
		if t.File != g.file.ProtoFile {
			gf, err := g.findGeneratedFile(t.File)
			if err != nil {
				panic("Can't find generated import" + err.Error())
			}
			if gf.OutGoPkg != g.file.OutGoPkg {
				return imports.New(gf.OutGoPkg) + "." + gf.Generator.gqlOutputTypeName(imports)(t)
			}
		}
		switch {
		case t.IsMessage():
			return g.file.GQLMessagePrefix + camelCaseSlice(t.Message.TypeName)
		case t.IsEnum():
			return g.file.GQLEnumsPrefix + camelCaseSlice(t.Enum.TypeName)
		case t.IsMap():
			return g.file.GQLMessagePrefix + camelCaseSlice(t.Map.Message.TypeName) + camelCase(t.Map.Field.Name) + "Map"
		}
		panic(t.String() + " is not handled in gqlOutputTypeName")
	}
}
func (g *gqlProtoDerivativeFileGenerator) gqlOutputTypeResolverResolver(imports *importer) (res func(t *parser.Type) string) {
	return func(t *parser.Type) string {
		if t.File != g.file.ProtoFile {
			gf, err := g.findGeneratedFile(t.File)
			if err != nil {
				panic("Can't find generated import" + err.Error())
			}
			if gf.OutGoPkg != g.file.OutGoPkg {
				return imports.New(gf.OutGoPkg) + "." + gf.Generator.gqlOutputTypeResolverResolver(imports)(t)
			}
		}
		switch {
		case t.Message != nil:
			return "Resolve" + camelCaseSlice(t.Message.TypeName)
		case t.Map != nil:
			return "Resolve" + camelCaseSlice(t.Map.Message.TypeName) + camelCase(t.Map.Field.Name) + "Map"
		}
		panic(t.String() + " is not handled in gqlOutputTypeResolverResolver")
	}
}
func (g *gqlProtoDerivativeFileGenerator) typeIsUsedInMessage(checkedMsgs parser.Messages, msg *parser.Message, typ *parser.Type) bool {
	if checkedMsgs.Contains(msg) {
		return false
	}
	for _, fld := range msg.Fields {
		if fld.Type == typ {
			return true
		}
		if fld.Type.IsMessage() && g.typeIsUsedInMessage(append(checkedMsgs.Copy(), msg), fld.Type.Message, typ) {
			return true
		}
	}
	for _, mfld := range msg.MapFields {
		if mfld.Map.ValueType == typ {
			return true
		}
		if mfld.Map.ValueType.IsMessage() && g.typeIsUsedInMessage(append(checkedMsgs.Copy(), msg), mfld.Map.ValueType.Message, typ) {
			return true
		}
	}
	for _, of := range msg.OneOffs {
		for _, fld := range of.Fields {
			if fld.Type == typ {
				return true
			}
			if fld.Type.IsMessage() && g.typeIsUsedInMessage(append(checkedMsgs.Copy(), msg), fld.Type.Message, typ) {
				return true
			}
		}
	}
	return false
}
func (g *gqlProtoDerivativeFileGenerator) needToGenerateTypeInputObject(typ *parser.Type) bool {
	for _, f := range g.generatedFiles {
		for _, s := range f.ProtoFile.Services {
			for _, mtd := range s.Methods {
				if g.typeIsUsedInMessage(nil, mtd.InputMessage, typ) {
					return true
				}
			}
		}
	}
	return false
}

func (g *gqlProtoDerivativeFileGenerator) needToGenerateTypeInputObjectResolver(typ *parser.Type) bool {
	for _, f := range g.generatedFiles {
		for _, s := range f.ProtoFile.Services {
			for _, mtd := range s.Methods {
				if mtd.InputMessage.Type == typ {
					return true
				}
				if g.typeIsUsedInMessage(nil, mtd.InputMessage, typ) {
					return true
				}
			}
		}
	}
	return false
}

func (g *gqlProtoDerivativeFileGenerator) needToGenerateTypeOutputObject(typ *parser.Type) (res bool) {
	for _, f := range g.generatedFiles {
		for _, s := range f.ProtoFile.Services {
			for _, mtd := range s.Methods {
				if mtd.OutputMessage.Type == typ {
					return true
				}
				if g.typeIsUsedInMessage(nil, mtd.OutputMessage, typ) {
					return true
				}
			}
		}
	}
	return false
}
func (g *gqlProtoDerivativeFileGenerator) methodIsQuery(m *parser.Method) bool {
	if srvCfg, ok := g.file.Services[m.Service.Name]; ok {
		if mtdCfg, ok := srvCfg.Methods[m.Name]; ok {
			return mtdCfg.RequestType == MethodTypeQuery
		}
	}
	return strings.HasPrefix(strings.ToLower(m.Name), "get")
}

func (g *gqlProtoDerivativeFileGenerator) serviceName(m *parser.Service) string {
	if srvCfg, ok := g.file.Services[m.Name]; ok && srvCfg.Alias != "" {
		return srvCfg.Alias
	}
	return m.Name
}
func (g *gqlProtoDerivativeFileGenerator) methodName(m *parser.Method) string {
	if srvCfg, ok := g.file.Services[m.Service.Name]; ok {
		if mtdCfg, ok := srvCfg.Methods[m.Name]; ok && mtdCfg.Alias != "" {
			return mtdCfg.Alias
		}
	}
	return m.Name
}
func (g *gqlProtoDerivativeFileGenerator) serviceHaveQueries(m *parser.Service) bool {
	for _, m := range m.Methods {
		if g.methodIsQuery(m) {
			return true
		}
	}
	return false
}

func (g *gqlProtoDerivativeFileGenerator) serviceHaveMutations(m *parser.Service) bool {
	for _, m := range m.Methods {
		if !g.methodIsQuery(m) {
			return true
		}
	}
	return false
}
func (g *gqlProtoDerivativeFileGenerator) messageHaveFieldsExceptError(msg *parser.Message) bool {
	ef := g.errorFieldOfMessage(msg)
	if ef == nil {
		return msg.HaveFields()
	}
	return msg.HaveFieldsExcept(ef.Name)
}
func (g *gqlProtoDerivativeFileGenerator) errorFieldOfMessage(msg *parser.Message) *ErrorField {
	m, ok := g.file.Messages[msg.Name]
	if !ok || m.ErrorField == "" {
		return nil
	}
	// Iterating over fields to be sure, that specified in config field exists
	for _, f := range msg.Fields {
		if f.Name == m.ErrorField {
			return &ErrorField{
				Name:     f.Name,
				Repeated: f.Repeated,
				Type:     f.Type,
			}
		}
	}
	for _, f := range msg.MapFields {
		if f.Name == m.ErrorField {
			return &ErrorField{
				Name:     f.Name,
				Repeated: false,
				Type:     f.Type,
			}
		}
	}
	for _, of := range msg.OneOffs {
		for _, f := range of.Fields {
			if f.Name == m.ErrorField {
				return &ErrorField{
					Name:     f.Name,
					Repeated: f.Repeated,
					Type:     f.Type,
				}
			}
		}
	}
	return nil
}
func (g *gqlProtoDerivativeFileGenerator) isErrorField(msg *parser.Message, name string) bool {
	if m, ok := g.file.Messages[msg.Name]; ok {
		return name == m.ErrorField
	}
	return false
}

func (g *gqlProtoDerivativeFileGenerator) fieldContextKey(msg *parser.Message, name string) string {
	return g.file.Messages[msg.Name].Fields[name].ContextKey
}

func (g *gqlProtoDerivativeFileGenerator) templateContext(imports *importer) map[string]interface{} {
	return map[string]interface{}{
		"File":            g.file.ProtoFile,
		"pkg":             g.file.OutGoPkgName,
		"protoPkg":        imports.New(g.file.GoProtoPkg),
		"gqlpkg":          imports.New(graphqlPkgPath),
		"interceptorspkg": imports.New(interceptorsPkgPath),
		"opentracingpkg":  imports.New(opentracingPkgPath),
		"tracerpkg":       imports.New(tracerPkg),
		"errorspkg":       imports.New("errors"),
		"ctxpkg":          imports.New("context"),
		"strconvpkg":      imports.New("strconv"),
		"debugpkg":        imports.New("runtime/debug"),
		"fmtpkg":          imports.New("fmt"),
		"ccase":           camelCase,
		"imports":         imports.Imports(),
		"tracerEnabled":   g.file.TracerEnabled,

		"FieldContextKey":                    g.fieldContextKey,
		"MessageErrorField":                  g.errorFieldOfMessage,
		"MessageHaveFieldsExceptError":       g.messageHaveFieldsExceptError,
		"IsErrorField":                       g.isErrorField,
		"MethodName":                         g.methodName,
		"ServiceName":                        g.serviceName,
		"MethodIsQuery":                      g.methodIsQuery,
		"ServiceHaveQueries":                 g.serviceHaveQueries,
		"ServiceHaveMutations":               g.serviceHaveMutations,
		"GoType":                             g.goTypeResolver(imports),
		"GQLInputTypeName":                   g.gqlInputTypeName(imports),
		"GQLInputTypeResolver":               g.gqlOutputTypeResolverResolver(imports),
		"GQLOutputTypeName":                  g.gqlOutputTypeName(imports),
		"NeedToGenerateTypeGQLInput":         g.needToGenerateTypeInputObject,
		"NeedToGenerateTypeGQLInputResolver": g.needToGenerateTypeInputObjectResolver,
		"NeedToGenerateTypeGQLOutput":        g.needToGenerateTypeOutputObject,
	}
}
func (g *gqlProtoDerivativeFileGenerator) findGeneratedFile(f *parser.File) (*gqlProtoDerivativeFile, error) {
	for _, fl := range g.generatedFiles {
		if fl.ProtoFile == f {
			return fl, nil
		}
	}
	return nil, errors.New("Not found")
}
func (g *gqlProtoDerivativeFileGenerator) goTypeResolver(imports *importer) (resolver func(t *parser.Type) string) {
	return func(t *parser.Type) string {
		switch {
		case t.Message != nil:
			var pkgPrefix string
			gf, err := g.findGeneratedFile(t.File)
			if err != nil {
				panic("Can't find generated import" + err.Error())
			}
			pkgPrefix = imports.New(gf.GoProtoPkg) + "."
			return pkgPrefix + snakeCamelCaseSlice(t.Message.TypeName)
		case t.Enum != nil:
			var pkgPrefix string
			gf, err := g.findGeneratedFile(t.File)
			if err != nil {
				panic("Can't find generated import" + err.Error())
			}
			pkgPrefix = imports.New(gf.GoProtoPkg) + "."
			return pkgPrefix + snakeCamelCaseSlice(t.Enum.TypeName)
		case t.Map != nil:
			kgt := resolver(t.Map.KeyType)
			vgt := resolver(t.Map.ValueType)
			res := "map[" + kgt + "]"
			if t.Map.ValueType.IsMessage() {
				res += "*"
			}
			res += vgt
			return res
		}
		res, err := g.scalarGoType(t.Scalar)
		if err != nil {
			panic(err) // as this function is template function, panic will be recovered
		}
		return res
	}
}

func (g *gqlProtoDerivativeFileGenerator) generate() error {
	tpl, err := template.New("template").Parse(bodyTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to parse template")
	}
	var imprts = new(importer)
	res := bytes.NewBuffer(nil)
	err = tpl.Execute(res, g.templateContext(imprts))
	if err != nil {
		return errors.Wrap(err, "failed to execute template")
	}
	headres := bytes.NewBuffer(nil)
	hdtpd, err := template.New("header").Parse(headTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to parse header template")
	}
	err = hdtpd.Execute(headres, g.templateContext(imprts))
	if err != nil {
		return errors.Wrap(err, "failed to execute header template")
	}
	r := bytes.Join([][]byte{headres.Bytes(), res.Bytes()}, nil)
	r, err = imports.Process(g.file.OutFilePath, r, &imports.Options{
		Comments: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to format generated code")
	}
	err = os.MkdirAll(g.file.OutDir, 0777)
	if err != nil {
		return errors.Wrapf(err, "failed to make output dir %s", g.file.OutDir)
	}
	err = ioutil.WriteFile(g.file.OutFilePath, r, 0600)
	return err
}

func newProtoGenerator(file *gqlProtoDerivativeFile, files []*gqlProtoDerivativeFile) *gqlProtoDerivativeFileGenerator {
	return &gqlProtoDerivativeFileGenerator{
		file:           file,
		generatedFiles: files,
	}
}
