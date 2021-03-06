package generator

import (
	"path/filepath"
	"strings"

	"github.com/saturn4er/proto2gql/parser"
)

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
