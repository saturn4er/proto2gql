package parser

import (
	"strings"

	"fmt"
	"github.com/emicklei/proto"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

var NotFoundErr = errors.New("not found")

type File struct {
	FilePath           string
	file               *proto.Proto
	PkgName            string
	importAliases      map[string]string
	paths              []string
	gqlPkg, scalarsPkg string
	Services           []*Service
	Messages           []*Message
	Maps               []*Map
	Enums              []*Enum
	Imports            []*File  // package name => File
	ParsedFiles        *[]*File // package name => File
}

func (f *File) FindMessage(pkg string, typeName []string) (*Message, error) {
	if pkg == f.PkgName {
		for _, msg := range f.Messages {
			if len(typeName) != len(msg.TypeName) {
				continue
			}
			if sameTypeNames(msg.TypeName, typeName) {
				return msg, nil
			}
		}
	}
	for _, imp := range f.Imports {
		if imp.PkgName == pkg {
			return imp.FindMessage(pkg, typeName)
		}
	}
	return nil, NotFoundErr
}
func (f *File) GetMap(mapField *proto.MapField) (*Map, error) {
	for _, mp := range f.Maps {
		if mp.Field == mapField {
			return mp, nil
		}
	}
	return nil, NotFoundErr
}
func (f *File) FindEnum(pkg string, typeName []string) (*Enum, error) {
	if pkg == f.PkgName {
		for _, enum := range f.Enums {
			if sameTypeNames(enum.TypeName, typeName) {
				return enum, nil
			}
		}
	}
	for _, imp := range f.Imports {
		if imp.PkgName == pkg {
			return imp.FindEnum(pkg, typeName)
		}
	}
	return nil, NotFoundErr
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
	TypeName      []string
	descriptor    *proto.Enum
}

type Method struct {
	Name          string
	QuotedComment string
	InputMessage  *Message
	OutputMessage *Message
	Service       *Service
}
type Message struct {
	Index           int
	Name            string
	QuotedComment   string
	Fields          []*Field
	MapFields       []*MapField
	OneOffs         []*OneOf
	GQLInputVarName string
	InputMessage    bool // Message used in fields of some service method input Message
	OutputMessage   bool // Message used for service method response
	Type            *ProtoType
	Descriptor      *proto.Message `json:"-"`
	TypeName        []string
	file            *File
	parentMsg       *Message
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

func (m *Message) String() string {
	return m.Name
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
	Index         int
	Name          string
	QuotedComment string
	Repeated      bool
	descriptor    *proto.Field `json:"-"`
	Type          *ProtoType
}

type MapField struct {
	Index         int
	Name          string
	QuotedComment string
	descriptor    *proto.MapField `json:"-"`
	Type          *ProtoType
}

type OneOf struct {
	Name   string
	Fields []*Field
}

func (f *File) resolveScalarType(t string, repeated bool) (*ProtoType, error) {
	if !typeIsScalar(t) {
		return nil, errors.Errorf("type %s is not scalar", t)
	}
	return &ProtoType{
		File:     f,
		Scalar:   t,
		Array:    repeated,
		Optional: true,
	}, nil
}
func (f *File) resolveMapValueType(msg *Message, valType string) (*ProtoType, error) {
	if typeIsScalar(valType) {
		return f.resolveScalarType(valType, false)
	}
	file, typename, t, err := f.findTypeInMessage(msg, valType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find type %s in message %s", t, msg.Name)
	}
	typ := &ProtoType{
		File: file,
	}
	switch t.(type) {
	case *proto.Message:
		m, err := f.FindMessage(file.PkgName, typename)
		if err != nil {
			return nil, errors.Errorf("can't find message (pkg: %s, TypeName: %v)", file.PkgName, typename)
		}
		typ.Message = m
	case *proto.Enum:
		enum, err := f.FindEnum(file.PkgName, typename)
		if err != nil {
			return nil, errors.Errorf("can't find enum (pkg: %s, TypeName: %v)", file.PkgName, typename)
		}
		typ.Enum = enum
	default:
		return nil, errors.Errorf("can't resolve map value type %s :( ", t)
	}
	return typ, nil
}

func (f *File) resolveFieldType(message *Message, field proto.Visitee, optional bool) (*ProtoType, error) {

	if mapField, ok := field.(*proto.MapField); ok {
		m, err := f.GetMap(mapField)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to find map")
		}
		return &ProtoType{
			File:     message.file,
			Optional: optional,
			Map:      m,
		}, nil
	}

	var fld *proto.Field
	var repeated bool
	switch f := field.(type) {
	case *proto.NormalField:
		fld = f.Field
		repeated = f.Repeated
	case *proto.OneOfField:
		fld = f.Field
	}
	switch fld.Type {
	case "double", "float", "int64", "uint64", "int32", "uint32", "fixed64", "fixed32", "bool", "string", "bytes", "sfixed32", "sfixed64", "sint32", "sint64":
		t, err := f.resolveScalarType(fld.Type, repeated)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve scalar type")
		}
		return t, nil

	}
	typ := &ProtoType{
		Optional: optional,
		Array:    repeated,
	}
	file, typename, t, err := f.findTypeInMessage(message, fld.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find Field %s type", fld.Name)
	}
	typ.File = file
	switch val := t.(type) {
	case *proto.Message:
		m, err := f.FindMessage(file.PkgName, typename)
		if err != nil {
			m, err = f.FindMessage(file.PkgName, typename)
			return nil, errors.Errorf("can't find message (pkg: %s, TypeName: %v)", file.PkgName, typename)
		}
		typ.Message = m
	case *proto.Enum:
		f.getTypename(val)
		enum, err := f.FindEnum(file.PkgName, typename)
		if err != nil {
			return nil, errors.Errorf("can't find enum (pkg: %s, TypeName: %v)", file.PkgName, typename)
		}
		typ.Enum = enum
	default:
		return nil, errors.Errorf("can't resolve ProtoType of Field %s:( ", fld.Name)
	}
	return typ, nil
}
func (f *File) messageByTypeName(typeName []string) (*Message, bool) {
	for _, msg := range f.Messages {
		if sameTypeNames(typeName, msg.TypeName) {
			return msg, true
		}
	}
	return nil, false
}
func (f *File) enumByTypeName(typeName []string) (*Enum, bool) {
	for _, e := range f.Enums {
		if sameTypeNames(typeName, e.TypeName) {
			return e, true
		}
	}
	return nil, false
}
func (f *File) parseAllUsedInServicesEntities() error {
	var handleMessage func(file *File, msg *proto.Message, isInput, isOutput bool, handled map[*proto.Message]bool) error

	handleMessage = func(file *File, msg *proto.Message, isInput, isOutput bool, hc map[*proto.Message]bool) error {
		if _, ok := hc[msg]; ok {
			return nil
		}
		typename, err := file.getTypename(msg)
		if err != nil {
			return errors.Wrap(err, "failed to get TypeName")
		}
		hc[msg] = true
		m, handled := file.messageByTypeName(typename)
		if !handled {
			m = &Message{
				file:          file,
				TypeName:      typename,
				Name:          msg.Name,
				Descriptor:    msg,
				QuotedComment: quoteComment(msg.Comment),
			}
			m.Type = &ProtoType{File: file, Message: m}
			file.Messages = append(file.Messages, m)
		}
		if isInput {
			m.InputMessage = true
		}
		if isOutput {
			m.OutputMessage = true
		}

		// handling submessages
		for _, el := range msg.Elements {
			switch fld := el.(type) {
			case *proto.MapField:
				if typeIsScalar(fld.Type) {
					break
				}
				file, typename, valTyp, err := file.findTypeInMessage(m, fld.Type)
				if err != nil {
					return errors.Wrapf(err, "failed to find Field %s map value type %s", fld.Name, fld.Type)
				}
				switch tv := valTyp.(type) {
				case *proto.Message:
					err = handleMessage(file, tv, isInput, isOutput, hc)
					if err != nil {
						return errors.Wrapf(err, "failed to handle type %s", fld.Type)
					}
				case *proto.Enum:
					if _, ok := file.enumByTypeName(typename); !ok {
						file.Enums = append(file.Enums, f.prepareEnum(file, tv, typename))
					}
				}
			case *proto.Oneof:
				for _, el := range fld.Elements {
					fld, ok := el.(*proto.OneOfField)
					if !ok {
						continue
					}
					if typeIsScalar(fld.Type) {
						continue
					}
					file, typename, typ, err := file.findTypeInMessage(m, fld.Type)
					if err != nil {
						return errors.Wrapf(err, "failed to find field %s type %s", fld.Name, fld.Type)
					}
					switch tv := typ.(type) {
					case *proto.Message:
						err = handleMessage(file, tv, isInput, isOutput, hc)

						if err != nil {
							return errors.Wrapf(err, "failed to message type %s", typename)
						}
					case *proto.Enum:
						if _, ok := file.enumByTypeName(typename); !ok {
							file.Enums = append(file.Enums, f.prepareEnum(file, tv, typename))
						}
					}
				}

			case *proto.NormalField:
				if typeIsScalar(fld.Type) {
					break
				}
				file, typename, typ, err := file.findTypeInMessage(m, fld.Type) // TODO: find Field type
				if err != nil {
					return errors.Wrapf(err, "failed to find type %s", fld.Type)
				}
				switch tv := typ.(type) {
				case *proto.Message:
					err = handleMessage(file, tv, isInput, isOutput, hc)
					if err != nil {
						return errors.Wrapf(err, "failed to handle message %s", tv.Name)
					}
				case *proto.Enum:
					if _, ok := file.enumByTypeName(typename); !ok {
						file.Enums = append(file.Enums, f.prepareEnum(file, tv, typename))
					}
				}
			}

		}
		return nil
	}
	var handledInput = make(map[*proto.Message]bool)
	var handledOutput = make(map[*proto.Message]bool)
	for _, el := range f.file.Elements {
		service, ok := el.(*proto.Service)
		if !ok {
			continue
		}
		for _, el := range service.Elements {
			method, ok := el.(*proto.RPC)
			if !ok {
				continue
			}
			file, _, inputMsg, err := f.findType(method.RequestType)
			if err != nil {
				return errors.Wrapf(err, "failed to find rpc request type %s", method.RequestType)
			}
			if err := handleMessage(file, inputMsg.(*proto.Message), true, false, handledInput); err != nil {
				return err // todo: wrap
			}
			file, _, outputMsg, err := f.findType(method.ReturnsType)
			if err != nil {
				return errors.Wrapf(err, "failed to find rpc response type %s", method.ReturnsType)
			}
			if err := handleMessage(file, outputMsg.(*proto.Message), false, true, handledOutput); err != nil {
				return err // todo: wrap
			}
		}

	}
	return nil
}

// return
func (f *File) parseMessagesFields() error {
	for _, msg := range f.Messages {
		if len(msg.Fields) > 0 {
			continue
		}
		for _, el := range msg.Descriptor.Elements {
			switch fld := el.(type) {
			case *proto.NormalField:
				typ, err := f.resolveFieldType(msg, fld, true)
				if err != nil {
					return errors.Wrapf(err, "failed to resolve Field %s.%s type", msg.Name, fld.Name)
				}
				f := &Field{
					Name:          fld.Name,
					QuotedComment: quoteComment(fld.Comment),
					descriptor:    fld.Field,
					Repeated:      fld.Repeated,
					Type:          typ,
				}
				msg.Fields = append(msg.Fields, f)
			case *proto.MapField:
				typ, err := f.resolveFieldType(msg, fld, true)
				if err != nil {
					return errors.Wrapf(err, "failed to resolve Field %s.%s type", msg.Name, fld.Name)
				}
				f := &MapField{
					Name:          fld.Name,
					QuotedComment: quoteComment(fld.Comment),
					descriptor:    fld,
					Type:          typ,
				}
				msg.MapFields = append(msg.MapFields, f)
			case *proto.Oneof:
				res := &OneOf{
					Name: fld.Name,
				}
				for _, el := range fld.Elements {
					switch fld := el.(type) {
					case *proto.OneOfField:
						typ, err := f.resolveFieldType(msg, fld, true)
						if err != nil {
							return errors.Wrapf(err, "failed to resolve Field %s.%s type", msg.Name, fld.Name)
						}
						f := &Field{
							Name:          fld.Name,
							QuotedComment: quoteComment(fld.Comment),
							descriptor:    fld.Field,
							Repeated:      false,
							Type:          typ,
						}
						res.Fields = append(res.Fields, f)
					default:
						fmt.Printf("Unknown oneof element : %T\n", el)
					}
				}
				msg.OneOffs = append(msg.OneOffs, res)
			}

		}
	}
	return nil
}

func (f *File) getTypename(v proto.Visitee) ([]string, error) {
	var res []string
	prependType := func(typ string) {
		res = append(res, "")
		copy(res[1:], res)
		res[0] = typ
	}
	switch v.(type) {
	case *proto.Message, *proto.Enum:
		vv := v
		for {
			if msg, ok := vv.(*proto.Message); ok {
				prependType(msg.Name)
				vv = msg.Parent
			} else if enum, ok := vv.(*proto.Enum); ok {
				prependType(enum.Name)
				vv = enum.Parent
			} else {
				return res, nil
			}
		}

	}
	return nil, errors.Errorf("can't get TypeName of %T", v)
}
func (f *File) findTypeInMessage(msg *Message, typ string) (*File, []string, proto.Visitee, error) {
	m := msg.Descriptor
	mels := m.Elements
	for {
		for _, el := range mels {
			switch v := el.(type) {
			case *proto.Message:
				if v.Name != typ {
					continue
				}
				typename, err := f.getTypename(v)
				if err != nil {
					return nil, nil, nil, errors.Wrapf(err, "failed to get TypeName of message %s", v.Name)
				}
				return msg.file, typename, v, nil
			case *proto.Enum:
				if v.Name != typ {
					continue
				}
				typename, err := f.getTypename(v)
				if err != nil {
					return nil, nil, nil, errors.Wrapf(err, "failed to get TypeName of message %s", v.Name)
				}
				return msg.file, typename, v, nil
			}
		}
		switch nm := m.Parent.(type) {
		case *proto.Message:
			m = nm
			mels = nm.Elements
			continue
		case *proto.Proto:
			return f.findType(typ)
		}
	}
}
func (f *File) findType(typ string) (*File, []string, proto.Visitee, error) {
	parts := strings.Split(typ, ".")
	if len(parts) == 1 {
		return f.TypeNamed(f.PkgName, []string{typ})
	}
	for i := 0; i < len(parts)-1; i++ {
		file, typename, res, err := f.TypeNamed(strings.Join(parts[:i+1], "."), parts[i+1:])
		if err != NotFoundErr {
			return file, typename, res, err
		}
	}
	return nil, nil, nil, NotFoundErr
}
func (f *File) TypeNamed(pkg string, typename []string) (*File, []string, proto.Visitee, error) {
	if pkg == "" || pkg == f.PkgName {
		var scanVisite func(el proto.Visitee, needle []string) proto.Visitee
		var handleEls func([]proto.Visitee, []string) proto.Visitee
		scanVisite = func(el proto.Visitee, needle []string) proto.Visitee {
			switch v := el.(type) {
			case *proto.Message:

				if v.Name == needle[0] {
					if len(needle) > 1 {
						return handleEls(v.Elements, needle[1:])
					}
					return v
				}
				return nil
			case *proto.Enum:
				if v.Name == needle[0] {
					if len(needle) > 1 {
						return nil
					}
					return v
				}
			}
			return nil
		}
		handleEls = func(els []proto.Visitee, needle []string) proto.Visitee {
			for _, el := range els {
				if res := scanVisite(el, needle); res != nil {
					return res
				}
			}
			return nil
		}
		if v := handleEls(f.file.Elements, typename); v != nil {
			tn, err := f.getTypename(v)
			if err != nil {
				return nil, nil, nil, errors.Wrap(err, "failed to resolve TypeName")
			}
			return f, tn, v, nil
		}

		for _, imp := range f.Imports {
			if imp.PkgName == "" || imp.PkgName == pkg {
				file, typename, typ, err := imp.TypeNamed(pkg, typename)
				if err != nil {
					if err == NotFoundErr {
						continue
					}
					return nil, nil, nil, err
				}
				return file, typename, typ, nil
			}
		}
		return nil, nil, nil, NotFoundErr
	}
	for _, imp := range f.Imports {
		if imp.PkgName == pkg {
			file, typename, typ, err := imp.TypeNamed(pkg, typename)
			if err != nil {
				if err == NotFoundErr {
					continue
				}
				return nil, nil, nil, err
			}
			return file, typename, typ, nil
		}
	}
	return nil, nil, nil, NotFoundErr
}
func (f *File) parseServices() error {
	for _, el := range f.file.Elements {
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
			file, typename, _, err := f.findType(method.RequestType)
			if err != nil {
				return errors.Errorf("can't find request message %s", method.RequestType)
			}
			imsg, err := f.FindMessage(file.PkgName, typename)
			if err != nil {
				return errors.Errorf("can't find parsed request message %s (pkg %s:, TypeName: %v)", method.RequestType, file.PkgName, typename)
			}
			file, typename, _, err = f.findType(method.ReturnsType)
			if err != nil {
				return errors.Errorf("can't find response message %s", method.ReturnsType)
			}
			omsg, err := f.FindMessage(file.PkgName, typename)
			if err != nil {
				return errors.Errorf("can't find parsed response message %s (pkg %s:, TypeName: %v)", method.ReturnsType, file.PkgName, typename)
			}
			mtd := Method{
				Name:          method.Name,
				QuotedComment: quoteComment(method.Comment),
				InputMessage:  imsg,
				OutputMessage: omsg,
				Service:       srv,
			}
			srv.Methods = append(srv.Methods, &mtd)
		}
		f.Services = append(f.Services, srv)
	}
	return nil
}
func (f *File) prepareEnum(file *File, enum *proto.Enum, typename []string) *Enum {
	e := &Enum{
		Name:          enum.Name,
		QuotedComment: quoteComment(enum.Comment),
		descriptor:    enum,
		TypeName:      typename,
		file:          file,
	}
	e.Type = &ProtoType{Enum: e, File: file}
	for _, el := range enum.Elements {
		val, ok := el.(*proto.EnumField)
		if !ok {
			continue
		}
		e.Values = append(e.Values, &EnumValue{
			Name:          val.Name,
			Value:         val.Integer,
			QuotedComment: quoteComment(val.Comment),
		})
	}
	return e
}

func (f *File) parseMaps() error {
	if len(f.Maps) > 0 {
		return nil
	}
	for _, msg := range f.Messages {
		for _, el := range msg.Descriptor.Elements {
			mapField, ok := el.(*proto.MapField)
			if !ok {
				continue
			}
			kt, err := f.resolveScalarType(mapField.KeyType, false)
			if err != nil {
				return errors.Wrap(err, "failed to resolve map key type")
			}
			vt, err := f.resolveMapValueType(msg, mapField.Type)
			if err != nil {
				return errors.Wrap(err, "failed to resolve map value type")
			}
			m := &Map{
				Message:   msg,
				KeyType:   kt,
				ValueType: vt,
				File:      f,
				Field:     mapField,
			}
			m.Type = &ProtoType{Map: m, File: f}
			f.Maps = append(f.Maps, m)
		}

	}
	return nil
}
func (f *File) findParsedFile(filePath string) (*File, error) {
	for _, file := range *f.ParsedFiles {
		if file.FilePath == filePath {
			return file, nil
		}
	}
	return nil, NotFoundErr
}
func (f *File) parseImports() error {
	for _, v := range f.file.Elements {
		imprt, ok := v.(*proto.Import)
		if !ok {
			continue
		}
		imprtPath, err := f.findImportFile(imprt.Filename)
		if err != nil {
			return err
		}
		absImprtPath, err := filepath.Abs(imprtPath)
		if err != nil {
			return errors.Wrap(err, "failed to resolve absolute import path")
		}
		pf, err := f.findParsedFile(absImprtPath)
		if err == nil {
			f.Imports = append(f.Imports, pf)
			continue
		} else if err != NotFoundErr {
			return err
		}

		importFile, err := ParseFile(f.ParsedFiles, f.importAliases, f.paths, imprtPath, false)
		if err != nil {
			return errors.Wrap(err, "can't parse import")
		}
		f.Imports = append(f.Imports, importFile)
		*f.ParsedFiles = append(*f.ParsedFiles, importFile)
	}
	return nil
}
func (f *File) findImportFile(filename string) (filePath string, err error) {
	if v, ok := f.importAliases[filename]; ok {
		filename = v
	}
	for _, path := range f.paths {
		p := filepath.Join(path, filename)
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	return "", errors.Errorf("can't find import %s in any of %s", filename, f.paths)
}
func (f *File) parseImportsEntities() error {
	for _, imprt := range f.Imports {
		err := imprt.parseAllUsedInServicesEntities()
		if err != nil {
			return errors.Wrap(err, "failed to parse messages")
		}
		err = imprt.parseMaps()
		if err != nil {
			return errors.Wrap(err, "failed to parse maps")
		}
		err = imprt.parseServices()
		if err != nil {
			return errors.Wrap(err, "failed to parse services")
		}
		err = imprt.parseMessagesFields()
		if err != nil {
			return errors.Wrap(err, "failed to parse messages fields")
		}
	}
	return nil
}
func ParseFile(parsedFiles *[]*File, aliases map[string]string, paths []string, filePath string, parseEntities bool) (*File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open File %s", filePath)
	}
	f, err := proto.NewParser(file).Parse()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse proto File %s", filePath)
	}
	var pkgName string
	for _, el := range f.Elements {
		pkg, ok := el.(*proto.Package)
		if ok {
			pkgName = pkg.Name
			break
		}
	}
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve absolute file apth")
	}
	res := &File{
		FilePath:      absPath,
		file:          f,
		PkgName:       pkgName,
		importAliases: aliases,
		paths:         paths,
		ParsedFiles:   parsedFiles,
	}
	err = res.parseImports()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse imports")
	}
	if parseEntities {
		err = res.parseAllUsedInServicesEntities()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse messages")
		}
		err = res.parseMaps() // TODO: think why we need it
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse maps")
		}
		err = res.parseServices()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse services")
		}
		err = res.parseMessagesFields()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse messages fields")
		}
		err = res.parseImportsEntities()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse imports entities")
		}
	}
	return res, nil
}
