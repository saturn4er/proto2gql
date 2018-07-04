package parser

import (
	"github.com/go-openapi/spec"
)

const (
	TypeUnknown  byte = iota
	TypeString
	TypeNumber
	TypeInt32
	TypeInt64
	TypeFloat32
	TypeFloat64
	TypeBoolean
	TypeArray
	TypeObject
	TypeMap
	TypeNull
)
const (
	ParameterPositionQuery  byte = iota
	ParameterPositionBody
	ParameterPositionParams
)

// Parsed DTO's
type Type struct {
	Type     byte
	Enum     []string
	Object   *Object
	ElemType *Type
}
type ObjectProperty struct {
	Name        string
	Description string
	Required    bool
	Type        Type
}
type Object struct {
	Name       string
	Route      []string
	Properties []ObjectProperty
}
type MethodParameter struct {
	Type        Type
	Position    byte
	Name        string
	Description string
	Required    bool
}
type MethodResponse struct {
	StatusCode  int
	Description string
	ResultType  *Type
}
type Tag struct {
	Name        string
	Description string
	Methods     []Method
}
type Method struct {
	Path             string
	OperationID      string
	Description      string
	HTTPMethod       string
	MethodParameters []MethodParameter
	Responses        []MethodResponse
}
type File struct {
	file     *spec.Swagger
	BasePath string
	Location string
	Tags     []Tag
	Objects  []Object
}
