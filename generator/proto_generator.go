package generator

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/parser"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const (
	scalarsPkgPath      = "github.com/saturn4er/proto2gql/api/scalars"
	interceptorsPkgPath = "github.com/saturn4er/proto2gql/api/interceptors"
	graphqlPkgPath      = "github.com/graphql-go/graphql"
	opentracingPkgPath  = "github.com/opentracing/opentracing-go"
	tracerPkg           = "github.com/saturn4er/proto2gql/api/tracer"
)

type protoGenerator struct {
	file           *generatedFile
	imports        importer
	generatedFiles []*generatedFile
	generatedNames *[]string
}

func (g *protoGenerator) gqlEnumVarName(i *parser.Enum) string {
	return g.file.GQLEnumsPrefix + i.Name
}
func (g *protoGenerator) scalarGoType(typ string) (string, error) {
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

func (g *protoGenerator) scalarGQLType(typ string) (string, error) {
	scalarsPkg := g.imports.New(scalarsPkgPath)
	gqlPkg := g.imports.New(graphqlPkgPath)
	switch typ {
	case "double":
		return scalarsPkg + ".GraphQLFloat64Scalar", nil
	case "float":
		return scalarsPkg + ".GraphQLFloat32Scalar", nil
	case "int64":
		return scalarsPkg + ".GraphQLInt64Scalar", nil
	case "uint64":
		return scalarsPkg + ".GraphQLUInt64Scalar", nil
	case "int32":
		return scalarsPkg + ".GraphQLInt32Scalar", nil
	case "uint32":
		return scalarsPkg + ".GraphQLUInt32Scalar", nil
	case "fixed64":
		return scalarsPkg + ".GraphQLUInt64Scalar", nil
	case "fixed32":
		return scalarsPkg + ".GraphQLUInt32Scalar", nil
	case "bool":
		return gqlPkg + ".Boolean", nil
	case "string":
		return gqlPkg + ".String", nil
	case "bytes":
		return scalarsPkg + ".GraphQLBytesScalar", nil
	case "sfixed32":
		return scalarsPkg + ".GraphQLInt32Scalar", nil
	case "sfixed64":
		return scalarsPkg + ".GraphQLInt64Scalar", nil
	case "sint32":
		return scalarsPkg + ".GraphQLInt32Scalar", nil
	case "sint64":
		return scalarsPkg + ".GraphQLInt64Scalar", nil
	}
	return "", errors.Errorf("%s is not scalar", typ)
}
func (g *protoGenerator) gqlInputTypeName(t *parser.ProtoType) string {
	if t.Message != nil && !t.Message.HaveFields() {
		scalarsPkg := g.imports.New(scalarsPkgPath)
		return scalarsPkg + ".NoDataScalar"
	}
	if t.IsScalar() {
		res, err := g.scalarGQLType(t.Scalar)
		if err != nil {
			panic(err.Error())
		}
		return res
	}
	if t.File != g.file.File {
		gf, err := g.findGeneratedFile(t.File)
		if err != nil {
			panic("Can't find generated import" + err.Error())
		}
		if gf.GoPkg != g.file.GoPkg {
			return g.imports.New(gf.GoPkg) + "." + gf.Generator.gqlInputTypeName(t)
		}
	}
	switch {
	case t.IsMessage():
		return g.file.GQLMessagePrefix + CamelCaseSlice(t.Message.TypeName) + "Input"
	case t.IsEnum():
		return g.file.GQLEnumsPrefix + CamelCaseSlice(t.Enum.TypeName)
	case t.IsMap():
		return g.file.GQLMessagePrefix + CamelCaseSlice(t.Map.Message.TypeName) + CamelCase(t.Map.Field.Name) + "MapInput"
	}
	panic(t.String() + " is not handled in gqlInputTypeName")
}
func (g *protoGenerator) gqlOutputTypeName(t *parser.ProtoType) string {
	if t.Message != nil && !g.messageHaveFieldsExceptError(t.Message) {
		scalarsPkg := g.imports.New(scalarsPkgPath)
		return scalarsPkg + ".NoDataScalar"
	}
	if t.IsScalar() {
		res, err := g.scalarGQLType(t.Scalar)
		if err != nil {
			panic(err.Error())
		}
		return res
	}
	if t.File != g.file.File {
		gf, err := g.findGeneratedFile(t.File)
		if err != nil {
			panic("Can't find generated import" + err.Error())
		}
		if gf.GoPkg != g.file.GoPkg {
			return g.imports.New(gf.GoPkg) + "." + gf.Generator.gqlOutputTypeName(t)
		}
	}
	switch {
	case t.IsMessage():
		return g.file.GQLMessagePrefix + CamelCaseSlice(t.Message.TypeName)
	case t.IsEnum():
		return g.file.GQLEnumsPrefix + CamelCaseSlice(t.Enum.TypeName)
	case t.IsMap():
		return g.file.GQLMessagePrefix + CamelCaseSlice(t.Map.Message.TypeName) + CamelCase(t.Map.Field.Name) + "Map"
	}
	panic(t.String() + " is not handled in gqlOutputTypeName")
}
func (g *protoGenerator) gqlOutputTypeResolverResolver(t *parser.ProtoType) string {
	if t.File != g.file.File {
		gf, err := g.findGeneratedFile(t.File)
		if err != nil {
			panic("Can't find generated import" + err.Error())
		}
		if gf.GoPkg != g.file.GoPkg {
			return g.imports.New(gf.GoPkg) + "." + gf.Generator.gqlOutputTypeResolverResolver(t)
		}
	}
	switch {
	case t.Message != nil:
		return "Resolve" + CamelCaseSlice(t.Message.TypeName)
	case t.Map != nil:
		return "Resolve" + CamelCaseSlice(t.Map.Message.TypeName) + CamelCase(t.Map.Field.Name) + "Map"
	}
	panic(t.String() + " is not handled in gqlOutputTypeResolverResolver")
}
func (g *protoGenerator) methodIsQuery(m *parser.Method) bool {
	if srvCfg, ok := g.file.Services[m.Service.Name]; ok {
		if mtdCfg, ok := srvCfg.Methods[m.Name]; ok {
			return mtdCfg.RequestType == MethodTypeQuery
		}
	}
	return strings.HasPrefix(strings.ToLower(m.Name), "get")
}

func (g *protoGenerator) serviceName(m *parser.Service) string {
	if srvCfg, ok := g.file.Services[m.Name]; ok {
		return srvCfg.Alias
	}
	return m.Name
}
func (g *protoGenerator) methodName(m *parser.Method) string {
	if srvCfg, ok := g.file.Services[m.Service.Name]; ok {
		if mtdCfg, ok := srvCfg.Methods[m.Name]; ok && mtdCfg.Alias != "" {
			return mtdCfg.Alias
		}
	}
	return m.Name
}
func (g *protoGenerator) serviceHaveQueries(m *parser.Service) bool {
	for _, m := range m.Methods {
		if g.methodIsQuery(m) {
			return true
		}
	}
	return false
}

func (g *protoGenerator) serviceHaveMutations(m *parser.Service) bool {
	for _, m := range m.Methods {
		if !g.methodIsQuery(m) {
			return true
		}
	}
	return false
}
func (g *protoGenerator) messageHaveFieldsExceptError(msg *parser.Message) bool {
	return msg.HaveFieldsExcept(g.messageErrorField(msg))
}
func (g *protoGenerator) messageErrorField(msg *parser.Message) string {
	m, ok := g.file.Messages[msg.Name]
	if !ok || m.ErrorField == "" {
		return ""
	}
	// Iterating over fields to be sure, that specified in config field exists
	for _, f := range msg.Fields {
		if f.Name == m.ErrorField {
			return f.Name
		}
	}
	for _, f := range msg.MapFields {
		if f.Name == m.ErrorField {
			return f.Name
		}
	}
	for _, of := range msg.OneOffs {
		for _, f := range of.Fields {
			if f.Name == m.ErrorField {
				return f.Name
			}
		}
	}
	return ""
}
func (g *protoGenerator) isErrorField(msg *parser.Message, name string) bool {
	if m, ok := g.file.Messages[msg.Name]; ok {
		return name == m.ErrorField
	}
	return false
}
func (g *protoGenerator) fieldContextKey(msg *parser.Message, name string) string {
	return g.file.Messages[msg.Name].Fields[name].ContextKey
}
func (g *protoGenerator) templateContext() map[string]interface{} {
	return map[string]interface{}{
		"File":            g.file.File,
		"pkg":             g.file.PkgName,
		"protoPkg":        g.imports.New(g.file.GoProtoPkg),
		"gqlpkg":          g.imports.New(graphqlPkgPath),
		"interceptorspkg": g.imports.New(interceptorsPkgPath),
		"opentracingpkg":  g.imports.New(opentracingPkgPath),
		"tracerpkg":       g.imports.New(tracerPkg),
		"errorspkg":       g.imports.New("errors"),
		"ctxpkg":          g.imports.New("context"),
		"strconvpkg":      g.imports.New("strconv"),
		"debugpkg":        g.imports.New("runtime/debug"),
		"fmtpkg":          g.imports.New("fmt"),
		"ccase":           parser.CamelCase,
		"imports":         g.imports.Imports(),
		"tracerEnabled":   g.file.TracerEnabled,

		"FieldContextKey":              g.fieldContextKey,
		"MessageErrorField":            g.messageErrorField,
		"MessageHaveFieldsExceptError": g.messageHaveFieldsExceptError,
		"IsErrorField":                 g.isErrorField,
		"MethodName":                   g.methodName,
		"ServiceName":                  g.serviceName,
		"MethodIsQuery":                g.methodIsQuery,
		"ServiceHaveQueries":           g.serviceHaveQueries,
		"ServiceHaveMutations":         g.serviceHaveMutations,
		"GoType":                       g.goTypeResolver,
		"GQLInputTypeName":             g.gqlInputTypeName,
		"GQLOutputTypeName":            g.gqlOutputTypeName,
		"GQLOutputTypeResolver":        g.gqlOutputTypeResolverResolver,
	}
}
func (g *protoGenerator) findGeneratedFile(f *parser.File) (*generatedFile, error) {
	for _, fl := range g.generatedFiles {
		if fl.File == f {
			return fl, nil
		}
	}
	return nil, errors.New("Not found")
}
func (g *protoGenerator) goTypeResolver(t *parser.ProtoType) string {
	switch {
	case t.Message != nil:
		var pkgPrefix string
		gf, err := g.findGeneratedFile(t.File)
		if err != nil {
			panic("Can't find generated import" + err.Error())
		}
		pkgPrefix = g.imports.New(gf.GoProtoPkg) + "."
		return pkgPrefix + SnakeCamelCaseSlice(t.Message.TypeName)
	case t.Enum != nil:
		var pkgPrefix string
		gf, err := g.findGeneratedFile(t.File)
		if err != nil {
			panic("Can't find generated import" + err.Error())
		}
		pkgPrefix = g.imports.New(gf.GoProtoPkg) + "."
		return pkgPrefix + SnakeCamelCaseSlice(t.Enum.TypeName)
	case t.Map != nil:
		kgt := g.goTypeResolver(t.Map.KeyType)
		vgt := g.goTypeResolver(t.Map.ValueType)
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

func (g *protoGenerator) generate() error {
	tpl, err := template.New("template").Parse(bodyTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to parse template")
	}
	res := bytes.NewBuffer(nil)
	err = tpl.Execute(res, g.templateContext())
	if err != nil {
		return errors.Wrap(err, "failed to execute template")
	}
	headres := bytes.NewBuffer(nil)
	hdtpd, err := template.New("header").Parse(headTemplate)
	if err != nil {
		panic(err)
	}
	err = hdtpd.Execute(headres, g.templateContext())
	if err != nil {
		panic(err)
	}
	r := bytes.Join([][]byte{headres.Bytes(), res.Bytes()}, nil)
	r, err = imports.Process(g.file.FilePath, r, &imports.Options{
		Comments: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to format generated code")
	}
	err = os.MkdirAll(g.file.Dir, 0777)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(g.file.FilePath, r, 0600)
	return err
}

func newProtoGenerator(file *generatedFile, files []*generatedFile) *protoGenerator {
	return &protoGenerator{
		file:           file,
		generatedFiles: files,
	}
}
