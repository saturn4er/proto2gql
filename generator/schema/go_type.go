package schema

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/saturn4er/proto2gql/generator/importer"
)

type GoType struct {
	Scalar    bool
	Kind      reflect.Kind
	ElemType  *GoType
	Elem2Type *GoType
	Name      string
	Pkg       string
}

func (g GoType) String(i *importer.Importer) string {
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

func typeIsScalar(p GoType) bool {
	switch p.Kind {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String:
		return true
	}
	return false
}
