package parser
//go:generate stringer -type=Kind
import (
	"github.com/go-openapi/spec"
)

type Kind byte

const (
	KindUnknown Kind = iota
	KindString
	KindInt32
	KindInt64
	KindFloat32
	KindFloat64
	KindBoolean
	KindArray
	KindObject
	KindMap
	KindFile
	KindNull
)
const (
	ParameterPositionQuery  byte = iota
	ParameterPositionBody
	ParameterPositionPath
	ParameterPositionHeader
	ParameterPositionFormData
)

var parameterPositions = map[string]byte{
	"path":     ParameterPositionPath,
	"query":    ParameterPositionQuery,
	"body":     ParameterPositionBody,
	"header":   ParameterPositionHeader,
	"formData": ParameterPositionFormData,
}

type Type interface {
	Kind() Kind
}
type Object struct {
	Name       string
	Route      []string
	Properties []ObjectProperty
}

func (Object) Kind() Kind {
	return KindObject
}

type Array struct {
	ElemType Type
}

func (Array) Kind() Kind {
	return KindArray
}

type Scalar struct {
	kind Kind
}

func (s Scalar) Kind() Kind {
	return s.kind
}

type Map struct {
	Route    []string
	ElemType Type
}

func (Map) Kind() Kind {
	return KindMap
}

type ObjectProperty struct {
	Name        string
	Description string
	Required    bool
	Type        Type
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
	ResultType  Type
}
type Tag struct {
	Name        string
	Description string
	Methods     []Method
}
type Method struct {
	Path        string
	OperationID string
	Description string
	HTTPMethod  string
	Parameters  []MethodParameter
	Responses   []MethodResponse
}
type File struct {
	file     *spec.Swagger
	BasePath string
	Location string
	Tags     []Tag
	Objects  []Object
}
