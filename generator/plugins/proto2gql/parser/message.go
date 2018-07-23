package parser

import "github.com/emicklei/proto"

type Messages []*Message

func (m Messages) Copy() Messages {
	result := make(Messages, len(m))
	copy(result, m)
	return result
}
func (m Messages) Contains(msg *Message) bool {
	for _, value := range m {
		if value == msg {
			return true
		}
	}
	return false
}

type Message struct {
	Name          string
	QuotedComment string
	Fields        []*Field
	MapFields     []*MapField
	OneOffs       []*OneOf
	Type          TypeInterface
	Descriptor    *proto.Message
	TypeName      TypeName
	File          *File
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
	Type          TypeInterface
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
	Type      TypeInterface
	Message   *Message
	KeyType   TypeInterface
	ValueType TypeInterface
	Field     *proto.MapField
	File      *File
}
