package common_generator

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"text/template"
	"reflect"
	"strings"
	"golang.org/x/tools/imports"
	"fmt"
	"os"
)

const (
	ScalarsPkgPath      = "github.com/saturn4er/proto2gql/api/scalars"
	InterceptorsPkgPath = "github.com/saturn4er/proto2gql/api/interceptors"
	GraphqlPkgPath      = "github.com/graphql-go/graphql"
	OpentracingPkgPath  = "github.com/opentracing/opentracing-go"
	TracerPkgPath       = "github.com/saturn4er/proto2gql/api/tracer"
	ErrorsPkgPath       = "github.com/pkg/errors"
)

type generator struct {
	Out  io.Writer
	File *File

	imports *Importer
}

func (g generator) importFunc(importPath string) func() string {
	return func() string {
		return g.imports.New(importPath)
	}
}
func (g generator) bodyTemplateContext() interface{} {
	return BodyContext{
		File:          g.File,
		Importer:      g.imports,
		TracerEnabled: true,
	}

}
func (g generator) goTypeStr(typ reflect.Type) string {
	if typeIsScalar(typ) {
		return typ.String()
	}
	switch typ.Kind() {
	case reflect.Slice:
		return "[]" + g.goTypeStr(typ.Elem())
	default:
		return g.imports.Prefix(typ.PkgPath()) + typ.Name()
	}
	panic("type " + typ.Kind().String() + " is not supported")
}

func (g generator) bodyTemplateFuncs() map[string]interface{} {
	return map[string]interface{}{
		"ctxPkg":          g.importFunc("context"),
		"debugPkg":        g.importFunc("runtime/debug"),
		"errorsPkg":       g.importFunc(ErrorsPkgPath),
		"gqlPkg":          g.importFunc(GraphqlPkgPath),
		"scalarsPkg":      g.importFunc(ScalarsPkgPath),
		"interceptorsPkg": g.importFunc(InterceptorsPkgPath),
		"opentracingPkg":  g.importFunc(OpentracingPkgPath),
		"tracerPkg":       g.importFunc(TracerPkgPath),
		"concat": func(st ...string) string {
			return strings.Join(st, "")
		},
		"isArray": func(typ reflect.Type) bool {
			return typ.Kind() == reflect.Slice
		},
		"goType": g.goTypeStr,
	}
}

func (g generator) headTemplateContext() map[string]interface{} {
	return map[string]interface{}{
		"imports": g.imports.Imports(),
		"package": g.File.PackageName,
	}

}
func (g generator) headTemplateFuncs() map[string]interface{} {
	return nil
}
func (g generator) generateBody() ([]byte, error) {
	buf := new(bytes.Buffer)
	bodyTpl, err := template.New("body").Funcs(g.bodyTemplateFuncs()).Parse(bodyTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}
	err = bodyTpl.Execute(buf, g.bodyTemplateContext())
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}
	return buf.Bytes(), nil
}

func (g generator) generateHead() ([]byte, error) {
	buf := new(bytes.Buffer)
	bodyTpl, err := template.New("head").Funcs(g.headTemplateFuncs()).Parse(headTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}
	err = bodyTpl.Execute(buf, g.headTemplateContext())
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}
	return buf.Bytes(), nil
}
func (g generator) generate() error {
	body, err := g.generateBody()
	if err != nil {
		return errors.Wrap(err, "failed to generate body")
	}
	head, err := g.generateHead()
	if err != nil {
		return errors.Wrap(err, "failed to generate head")
	}
	r := bytes.Join([][]byte{
		head,
		body,
	}, nil)

	res, err := imports.Process("file", r, &imports.Options{
		Comments: true,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		r = res
	}
	_, err = g.Out.Write(r)
	if err != nil {
		return errors.Wrap(err, "failed to write  output")
	}
	return nil
}

func Generate(file *File, w io.Writer) error {
	g := generator{
		File: file,
		Out:  w,

		imports: &Importer{
			CurrentPackage: file.PackagePath,
		},
	}
	return g.generate()
}
