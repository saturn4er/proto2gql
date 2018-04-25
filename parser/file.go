package parser

import (
	"strings"

	"github.com/emicklei/proto"
	"github.com/pkg/errors"
)

type File struct {
	FilePath           string
	protoFile          *proto.Proto
	PkgName            string
	importAliases      map[string]string
	paths              []string
	gqlPkg, scalarsPkg string
	Services           []*Service
	Messages           []*Message
	Enums              []*Enum
	Imports            []*File
	ParsedFiles        *[]*File
}

type Service struct {
	Name          string
	QuotedComment string
	Methods       []*Method
}
type EnumValue struct {
	Name          string
	Value         int
	QuotedComment string
}
type Enum struct {
	Name          string
	QuotedComment string
	Values        []*EnumValue
	Type          *ProtoType
	file          *File
	TypeName      TypeName
	Descriptor    *proto.Enum
}

type Method struct {
	Name          string
	QuotedComment string
	InputMessage  *Message
	OutputMessage *Message
	Service       *Service
}
type Message struct {
	Name          string
	QuotedComment string
	Fields        []*Field
	MapFields     []*MapField
	OneOffs       []*OneOf
	Type          *ProtoType
	Descriptor    *proto.Message `json:"-"`
	TypeName      TypeName
	file          *File
	parentMsg     *Message
}

func (m *Message) HaveFields() bool {
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
func (m *Message) HaveFieldsExcept(field string) bool {
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

type Map struct {
	Type      *ProtoType
	Message   *Message
	KeyType   *ProtoType
	ValueType *ProtoType
	Field     *proto.MapField
	File      *File
}

type Field struct {
	Name          string
	QuotedComment string
	Repeated      bool
	descriptor    *proto.Field `json:"-"`
	Type          *ProtoType
}

type MapField struct {
	Name          string
	QuotedComment string
	descriptor    *proto.MapField `json:"-"`
	Type          *ProtoType
	Map           *Map
}

type OneOf struct {
	Name   string
	Fields []*Field
}

func (f *File) MessageByTypeName(typeName TypeName) (*Message, bool) {
	for _, msg := range f.Messages {
		if msg.TypeName.Equal(typeName) {
			return msg, true
		}
	}
	return nil, false
}
func (f *File) EnumByTypeName(typeName TypeName) (*Enum, bool) {
	for _, e := range f.Enums {
		if e.TypeName.Equal(typeName) {
			return e, true
		}
	}
	return nil, false
}
func (f *File) FindTypeInMessage(msg *Message, typ string) (*ProtoType, bool) {
	if typeIsScalar(typ) {
		return &ProtoType{Scalar: typ, File: f}, true
	}
	ms, ok := f.MessageByTypeName(msg.TypeName.NewSubTypeName(typ))
	if ok {
		return ms.Type, true
	}

	enum, ok := f.EnumByTypeName(msg.TypeName.NewSubTypeName(typ))
	if ok {
		return enum.Type, true
	}
	if msg.parentMsg != nil {
		return f.FindTypeInMessage(msg.parentMsg, typ)
	}
	return f.FindType(typ)
}

func (f *File) FindType(typ string) (*ProtoType, bool) {
	if typeIsScalar(typ) {
		return &ProtoType{Scalar: typ, File: f}, true
	}
	parts := strings.Split(typ, ".")
	msg, ok := f.MessageByTypeName(parts)
	if ok {
		return msg.Type, true
	}
	en, ok := f.EnumByTypeName(parts)
	if ok {
		return en.Type, true
	}
	for _, imp := range f.Imports {
		if imp.PkgName == f.PkgName {
			it, ok := imp.FindType(typ)
			if ok {
				return it, true
			}
		}
	}
	for i := 0; i < len(parts)-1; i++ {
		pkg, typ := strings.Join(parts[:i+1], "."), strings.Join(parts[i+1:], ".")
		for _, imp := range f.Imports {
			if imp.PkgName == pkg {
				return imp.FindType(typ)
			}
		}
	}
	return nil, false
}
func (f *File) parseServices() error {
	for _, el := range f.protoFile.Elements {
		service, ok := el.(*proto.Service)
		if !ok {
			continue
		}
		srv := &Service{
			Name:          service.Name,
			QuotedComment: quoteComment(service.Comment),
		}
		for _, el := range service.Elements {
			method, ok := el.(*proto.RPC)
			if !ok {
				continue
			}
			reqTyp, ok := f.FindType(method.RequestType)
			if !ok {
				return errors.Errorf("can't find request message %s", method.RequestType)
			}
			retTyp, ok := f.FindType(method.ReturnsType)
			if !ok {
				return errors.Errorf("can't find request message %s", method.RequestType)
			}
			mtd := &Method{
				Name:          method.Name,
				QuotedComment: quoteComment(method.Comment),
				InputMessage:  reqTyp.Message,
				OutputMessage: retTyp.Message,
				Service:       srv,
			}
			srv.Methods = append(srv.Methods, mtd)
		}
		f.Services = append(f.Services, srv)
	}
	return nil
}
func (f *File) ParseMessagesFields() error {
	for _, msg := range f.Messages {
		for _, el := range msg.Descriptor.Elements {
			switch fld := el.(type) {
			case *proto.NormalField:
				typ, ok := f.FindTypeInMessage(msg, fld.Type)
				if !ok {
					return errors.Errorf("failed to find message %s field %s type", strings.Join(msg.TypeName, "."), fld.Name)
				}
				fl := &Field{
					Name:          fld.Name,
					QuotedComment: quoteComment(fld.Comment),
					Repeated:      fld.Repeated,
					descriptor:    fld.Field,
					Type:          typ,
				}
				msg.Fields = append(msg.Fields, fl)
			case *proto.MapField:
				ktyp, ok := f.FindTypeInMessage(msg, fld.KeyType)
				if !ok {
					return errors.Errorf("failed to find message %s field %s type", strings.Join(msg.TypeName, "."), fld.Name)
				}
				vtyp, ok := f.FindTypeInMessage(msg, fld.Type)
				if !ok {
					return errors.Errorf("failed to find message %s field %s type", strings.Join(msg.TypeName, "."), fld.Name)
				}
				mp := &Map{
					Message:   msg,
					KeyType:   ktyp,
					ValueType: vtyp,
					Field:     fld,
					File:      f,
				}
				t := &ProtoType{Map: mp, File: f}
				mp.Type = t
				mf := &MapField{
					Name:          fld.Name,
					QuotedComment: quoteComment(fld.Comment),
					descriptor:    fld,
					Type:          t,
					Map:           mp,
				}
				msg.MapFields = append(msg.MapFields, mf)
			case *proto.Oneof:
				of := &OneOf{
					Name: fld.Name,
				}
				for _, el := range fld.Elements {
					fld, ok := el.(*proto.OneOfField)
					if !ok {
						continue
					}
					typ, ok := f.FindTypeInMessage(msg, fld.Type)
					if !ok {
						return errors.Errorf("failed to find message %s field %s type", strings.Join(msg.TypeName, "."), fld.Name)
					}
					of.Fields = append(of.Fields, &Field{
						Name:          fld.Name,
						QuotedComment: quoteComment(fld.Comment),
						Repeated:      false,
						descriptor:    fld.Field,
						Type:          typ,
					})
				}
				msg.OneOffs = append(msg.OneOffs, of)
			}
		}
	}
	return nil
}
func (f *File) parseMessages() {
	for _, el := range f.protoFile.Elements {
		msg, ok := el.(*proto.Message)
		if !ok {
			continue
		}
		m := message(f, msg, TypeName{msg.Name}, nil)
		f.Messages = append(f.Messages, m)
		f.parseMessagesInMessage(TypeName{msg.Name}, m)
	}
}
func (f *File) parseMessagesInMessage(msgTypeName TypeName, msg *Message) {
	for _, el := range msg.Descriptor.Elements {
		switch elv := el.(type) {
		case *proto.Message:
			tn := msgTypeName.NewSubTypeName(elv.Name)
			m := message(f, elv, tn, msg)
			f.Messages = append(f.Messages, m)
			f.parseMessagesInMessage(tn, m)
		}
	}
}

func (f *File) parseEnums() {
	for _, el := range f.protoFile.Elements {
		switch val := el.(type) {
		case *proto.Enum:
			f.Enums = append(f.Enums, enum(f, val, TypeName{val.Name}))
		case *proto.Message:
			f.parseEnumsInMessage(TypeName{val.Name}, val)
		}

	}
}

func (f *File) parseEnumsInMessage(msgTypeName TypeName, msg *proto.Message) {
	for _, el := range msg.Elements {
		switch elv := el.(type) {
		case *proto.Message:
			f.parseEnumsInMessage(msgTypeName.NewSubTypeName(elv.Name), elv)
		case *proto.Enum:
			f.Enums = append(f.Enums, enum(f, elv, msgTypeName.NewSubTypeName(elv.Name)))
		}
	}
}
