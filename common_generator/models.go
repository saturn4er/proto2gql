package common_generator

import (
	"reflect"
)

type TypeResolver func(ctx BodyContext) string
type ValueResolver func(arg string, ctx BodyContext) string

type InputObjectResolver struct {
	FunctionName string
	OutputGoType reflect.Type
	Fields       []InputObjectResolverField
}

type InputObjectResolverField struct {
	Name              string
	GraphqlInputField string
	GoType            reflect.Type
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
	GoObjectType  reflect.Type
	GoObjectField string
	NeedCast      bool
	CastTo        reflect.Type
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
	KeyGoType              reflect.Type
	ValueGoType            reflect.Type
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
	KeyGoType       reflect.Type
	ValueGoType     reflect.Type
}
type Service struct {
	Name string
}
type File struct {
	PackageName             string
	PackagePath             string
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
