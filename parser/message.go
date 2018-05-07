package parser

import "github.com/emicklei/proto"

type Message struct {
	Name          string
	QuotedComment string
	Fields        []*Field
	MapFields     []*MapField
	OneOffs       []*OneOf
	Type          *Type
	Descriptor    *proto.Message
	TypeName      TypeName
	file          *File
	parentMsg     *Message
}

func (m Message) HaveFields() bool {
	if len(m.Fields) > 0 || len(m.MapFields) > 0 {
		return true
	}
	for _, of := range m.OneOffs {
		if len(of.Fields) > 0 {
			return true
		}
	}
	return false
}
func (m Message) HaveFieldsExcept(field string) bool {
	for _, f := range m.Fields {
		if f.Name != field {
			return true
		}
	}
	for _, f := range m.MapFields {
		if f.Name != field {
			return true
		}
	}
	for _, of := range m.OneOffs {
		for _, f := range of.Fields {
			if f.Name != field {
				return true
			}
		}
	}
	return false
}

type Field struct {
	Name          string
	QuotedComment string
	Repeated      bool
	descriptor    *proto.Field
	Type          *Type
}

type MapField struct {
	Name          string
	QuotedComment string
	descriptor    *proto.MapField
	Type          *Type
	Map           *Map
}

type OneOf struct {
	Name   string
	Fields []*Field
}

type Map struct {
	Type      *Type
	Message   *Message
	KeyType   *Type
	ValueType *Type
	Field     *proto.MapField
	File      *File
}
