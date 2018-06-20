package proto

import (
	"go/build"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/common"
	"github.com/saturn4er/proto2gql/generator/proto/parser"
)

var goTypesScalars = map[string]common.GoType{
	"double": {Scalar: true, Kind: reflect.Float64},
	"float":  {Scalar: true, Kind: reflect.Float32},
	"bool":   {Scalar: true, Kind: reflect.Bool},
	"string": {Scalar: true, Kind: reflect.String},

	"int64":    {Scalar: true, Kind: reflect.Int64},
	"sfixed64": {Scalar: true, Kind: reflect.Int64},
	"sint64":   {Scalar: true, Kind: reflect.Int64},

	"int32":    {Scalar: true, Kind: reflect.Int32},
	"sfixed32": {Scalar: true, Kind: reflect.Int32},
	"sint32":   {Scalar: true, Kind: reflect.Int32},

	"uint32":  {Scalar: true, Kind: reflect.Uint32},
	"fixed32": {Scalar: true, Kind: reflect.Uint32},

	"uint64":  {Scalar: true, Kind: reflect.Uint64},
	"fixed64": {Scalar: true, Kind: reflect.Uint64},
}

func goTypeByParserType(typ *parser.Type) (_ common.GoType, err error) {
	if typ.IsScalar() {
		res, ok := goTypesScalars[typ.Scalar]
		if !ok {
			err = errors.New("unknown scalar")
			return
		}
		return res, nil
	}
	if typ.IsEnum() {
		return common.GoType{
			Pkg:  typ.File.GoPackage,
			Name: typ.Enum.Name,
			Kind: reflect.Int32,
		}, nil
	}
	if typ.IsMessage() {
		msgType := &common.GoType{
			Pkg:  typ.File.GoPackage,
			Name: snakeCamelCaseSlice(typ.Message.TypeName),
			Kind: reflect.Struct,
		}
		return common.GoType{
			Pkg:      typ.File.GoPackage,
			Kind:     reflect.Ptr,
			ElemType: msgType,
		}, nil
	}
	err = errors.Errorf("unknown type " + typ.String())
	return
}
func resolveGoPackage(path, vendorPath string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", errors.Wrap(err, "failed to resolve absolute filepath")
	}
	var prefixes []string
	if vendorPath != "" {
		absVendorPath, err := filepath.Abs(vendorPath)
		if err != nil {
			return "", errors.Wrap(err, "failed to resolve absolute vendor path")
		}
		prefixes = append(prefixes, absVendorPath)
	}
	absGoPath, err := filepath.Abs(build.Default.GOPATH)
	if err != nil {
		return "", errors.Wrap(err, "failed to resolve absolute gopath")
	}
	prefixes = append(prefixes, filepath.Join(absGoPath, "src"))

	for _, prefix := range prefixes {
		if strings.HasPrefix(path, prefix) {
			return strings.TrimLeft(strings.TrimPrefix(path, prefix), " "+string(filepath.Separator)), nil
		}
	}
	return "", errors.Errorf("path '%s' is outside GOPATH or Vendor folder", path)
}

func mergePathsConfig(in ...[]string) []string {
	var res []string
	for i := len(in) - 1; i >= 0; i-- {
		res = append(res, in[i]...)
	}
	return res
}
func mergeStringsConfig(in ...string) string {
	var l int
	for i, value := range in {
		if len(value) > 0 {
			l = i
		}
	}
	return in[l]
}
func mergeAliases(in ...map[string]string) map[string]string {
	var res = make(map[string]string)
	for _, i := range in {
		for k, v := range i {
			res[k] = v
		}
	}
	return res
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func camelCase(s string) string {
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

// camelCaseSlice is like camelCase, but the argument is a slice of strings to
// be joined with "_".
func camelCaseSlice(elem []string) string      { return camelCase(strings.Join(elem, "")) }
func snakeCamelCaseSlice(elem []string) string { return camelCase(strings.Join(elem, "_")) }

func isSamePackage(f1, f2 *parser.File) bool {
	return f1.PkgName == f2.PkgName && filepath.Dir(f1.FilePath) == filepath.Dir(f2.FilePath)
}
