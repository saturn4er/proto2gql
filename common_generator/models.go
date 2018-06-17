package common_generator

import (
	"reflect"
)

type TypeResolver func(ctx BodyContext) string

type InputObjectResolver struct {
	FunctionName string
	OutputGoType string
	Fields       []InputObjectResolverField
}
type InputObjectResolverField struct {
	Name              string
	GraphqlInputField string
	GoType            reflect.Type
}

type InputObject struct {
	VariableName string
	GraphQLName  string
	Fields       []ObjectField
}
type ObjectField struct {
	Name string
	Type TypeResolver
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
type File struct {
	Package              string
	Enums                []Enum
	InputObjects         []InputObject
	InputObjectResolvers []InputObjectResolver
}

type BodyContext struct {
	File          *File
	Importer      *Importer
	TracerEnabled bool
}
