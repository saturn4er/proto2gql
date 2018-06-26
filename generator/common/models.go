package common

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type TypeResolver func(ctx BodyContext) string
type ValueResolver func(arg string, ctx BodyContext) string
type AssigningWrapper func(arg string, ctx BodyContext) string
type PayloadErrorChecker func(arg string) string
type PayloadErrorAccessor func(arg string) string

type GoType struct {
	Scalar    bool
	Kind      reflect.Kind
	ElemType  *GoType
	Elem2Type *GoType
	Name      string
	Pkg       string
}

func (g GoType) String(i *Importer) string {
	if typeIsScalar(g) && g.Name == "" {
		return g.Kind.String()
	}
	switch g.Kind {
	case reflect.Slice:
		return "[]" + g.ElemType.String(i)
	case reflect.Ptr:
		return "*" + g.ElemType.String(i)
	case reflect.Struct, reflect.Interface:
		return i.Prefix(g.Pkg) + g.Name
	case reflect.Map:
		return "map[" + g.ElemType.String(i) + "]" + g.Elem2Type.String(i)
	}
	if g.Name != "" {
		return i.Prefix(g.Pkg) + g.Name
	}
	spew.Dump(g)
	panic("type " + g.Kind.String() + " is not supported")
}

type InputObjectResolver struct {
	FunctionName string
	OutputGoType GoType
	Fields       []InputObjectResolverField
	OneOfFields  []InputObjectResolverOneOf
}
type InputObjectResolverOneOf struct {
	OutputFieldName string
	Fields          []InputObjectResolverOneOfField
}
type InputObjectResolverOneOfField struct {
	GraphQLInputFieldName string
	ValueResolver         ValueResolver
	ResolverWithError     bool
	AssigningWrapper      AssigningWrapper
}
type InputObjectResolverField struct {
	OutputFieldName       string
	GraphQLInputFieldName string
	GoType                GoType
	ValueResolver         ValueResolver
	ResolverWithError     bool
}

type InputObject struct {
	VariableName string
	GraphQLName  string
	Fields       []ObjectField
}
type ObjectField struct {
	Name          string
	Type          TypeResolver
	GoObjectField string
	NeedCast      bool
	CastTo        GoType
}
type OutputObject struct {
	VariableName string
	GraphQLName  string
	GoType       GoType
	Fields       []ObjectField
}
type Enum struct {
	VariableName string
	GraphQLName  string
	Comment      string
	Values       []EnumValue
}
type EnumValue struct {
	Name    string
	Value   int
	Comment string
}
type MapInputObject struct {
	VariableName    string
	GraphQLName     string
	KeyObjectType   TypeResolver
	ValueObjectType TypeResolver
}
type MapInputObjectResolver struct {
	FunctionName           string
	KeyGoType              GoType
	ValueGoType            GoType
	KeyResolver            ValueResolver
	KeyResolverWithError   bool
	ValueResolver          ValueResolver
	ValueResolverWithError bool
}
type MapOutputObject struct {
	VariableName    string
	GraphQLName     string
	KeyObjectType   TypeResolver
	ValueObjectType TypeResolver
}
type Service struct {
	Name          string
	CallInterface GoType
	Methods       []Method
}
type Method struct {
	Name                   string
	GraphQLOutputType      TypeResolver
	Arguments              []MethodArgument
	RequestResolver        ValueResolver
	RequestResolverWithErr bool
	CallMethod             string
	RequestType            GoType
	PayloadErrorChecker    PayloadErrorChecker
	PayloadErrorAccessor   PayloadErrorAccessor
}
type MethodArgument struct {
	Name string
	Type TypeResolver
}
type File struct {
	PackageName             string
	Package                 string
	Enums                   []Enum
	OutputObjects           []OutputObject
	InputObjects            []InputObject
	InputObjectResolvers    []InputObjectResolver
	MapInputObjects         []MapInputObject
	MapInputObjectResolvers []MapInputObjectResolver
	MapOutputObjects        []MapOutputObject
	Services                []Service
}

type BodyContext struct {
	File          *File
	Importer      *Importer
	TracerEnabled bool
}
