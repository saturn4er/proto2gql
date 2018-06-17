package common_generator

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"text/template"
)

const (
	ScalarsPkgPath      = "github.com/saturn4er/proto2gql/api/scalars"
	InterceptorsPkgPath = "github.com/saturn4er/proto2gql/api/interceptors"
	GraphqlPkgPath      = "github.com/graphql-go/graphql"
	OpentracingPkgPath  = "github.com/opentracing/opentracing-go"
	TracerPkgPath       = "github.com/saturn4er/proto2gql/api/tracer"
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
func (g generator) bodyTemplateFuncs() map[string]interface{} {
	return map[string]interface{}{
		"ctxPkg":          g.importFunc("context"),
		"debugPkg":        g.importFunc("runtime/debug"),
		"gqlPkg":          g.importFunc(GraphqlPkgPath),
		"scalarsPkg":      g.importFunc(ScalarsPkgPath),
		"interceptorsPkg": g.importFunc(InterceptorsPkgPath),
		"opentracingPkg":  g.importFunc(OpentracingPkgPath),
		"tracerPkg":       g.importFunc(TracerPkgPath),
	}
}

func (g generator) headTemplateContext() map[string]interface{} {
	return map[string]interface{}{
		"imports": g.imports.Imports(),
		"package": g.File.Package,
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
	res := bytes.Join([][]byte{
		head,
		body,
	}, nil)
	_, err = g.Out.Write(res)
	if err != nil {
		return errors.Wrap(err, "failed to write  output")
	}
	return nil
}

func Generate(file *File, w io.Writer) error {
	g := generator{
		File: file,
		Out:  w,

		imports: new(Importer),
	}
	return g.generate()
}
