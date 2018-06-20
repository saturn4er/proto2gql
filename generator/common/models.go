package common

import (
	"reflect"
)

type TypeResolver func(ctx BodyContext) string
type ValueResolver func(arg string, ctx BodyContext) string

type GoType struct {
	Scalar   bool
	Kind     reflect.Kind
	ElemType *GoType
	Name     string
	Pkg      string
}

type InputObjectResolver struct {
	FunctionName string
	OutputGoType GoType
	Fields       []InputObjectResolverField
}

type InputObjectResolverField struct {
	Name              string
	GraphqlInputField string
	GoType            GoType
	ValueResolver     ValueResolver
	ResolverWithError bool
}

type InputObject struct {
	VariableName string
	GraphQLName  string
	Fields       []ObjectField
}
type ObjectField struct {
	Name          string
	Type          TypeResolver
	GoObjectType  GoType
	GoObjectField string
	NeedCast      bool
	CastTo        GoType
}
type OutputObject struct {
	VariableName string
	GraphQLName  string
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
	KeyGoType       GoType
	ValueGoType     GoType
}
type Service struct {
	Name          string
	CallInterface GoType
	Methods       []Method
}
type Method struct {
	Name                        string
	ResultType                  TypeResolver
	Arguments                   []MethodArguments
	RequestResolverFunctionName string
	CallMethod                  string
	RequestType                 GoType
	ResponseType                GoType
}
type MethodArguments struct {
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
