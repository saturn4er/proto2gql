package parser

import "github.com/emicklei/proto"

type Enum struct {
	Name          string
	QuotedComment string
	Values        []*EnumValue
	Type          *Type
	File          *File
	TypeName      TypeName
	Descriptor    *proto.Enum
}

type EnumValue struct {
	Name          string
	Value         int
	QuotedComment string
}

func newEnum(file *File, enum *proto.Enum, typeName []string) *Enum {
	m := &Enum{
		Name:          enum.Name,
		QuotedComment: quoteComment(enum.Comment),
		Descriptor:    enum,
		Type:          &Type{File: file},
		TypeName:      typeName,
		File:          file,
	}
	m.Type.Enum = m
	for _, v := range enum.Elements {
		value, ok := v.(*proto.EnumField)
		if !ok {
			continue
		}
		m.Values = append(m.Values, &EnumValue{
			Name:          value.Name,
			Value:         value.Integer,
			QuotedComment: quoteComment(value.Comment),
		})
	}
	return m
}
