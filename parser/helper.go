package parser

import (
	"github.com/emicklei/proto"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func typeIsScalar(typ string) bool {
	switch typ {
	case "double", "float", "int32", "int64", "uint32", "uint64", "sint32", "sint64", "fixed32", "fixed64", "sfixed32", "sfixed64", "bool", "string", "bytes":
		return true
	}
	return false
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

func sameTypeNames(tn1 []string, tn2 []string) bool {
	if len(tn1) != len(tn2) {
		return false
	}
	for i, tn := range tn1 {
		if tn != tn2[i] {
			return false
		}
	}
	return true
}

func quoteComment(comments *proto.Comment) string {
	if comments == nil {
		return `""`
	}
	return strconv.Quote(strings.TrimSpace(strings.Join(comments.Lines, "\n")))
}

func getVisiteeTypename(v proto.Visitee) ([]string, error) {
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

func resolveFilePkgName(file *proto.Proto) string {
	for _, el := range file.Elements {
		if p, ok := el.(*proto.Package); ok {
			return p.Name
		}
	}
	return ""
}

func message(file *File, msg *proto.Message, typeName []string, parent *Message) *Message {
	m := &Message{
		Name:          msg.Name,
		QuotedComment: quoteComment(msg.Comment),
		Descriptor:    msg,
		Type:          &ProtoType{File: file},
		TypeName:      typeName,
		file:          file,
		parentMsg:     parent,
	}
	m.Type.Message = m
	return m
}

func enum(file *File, enum *proto.Enum, typeName []string) *Enum {
	m := &Enum{
		Name:          enum.Name,
		QuotedComment: quoteComment(enum.Comment),
		Descriptor:    enum,
		Type:          &ProtoType{File: file},
		TypeName:      typeName,
		file:          file,
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
