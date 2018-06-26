package schema

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

type ConstructorResolver func() string

type SchemaNodeConfig struct {
	Type           string             `yaml:"type"` // "OBJECT|SERVICE"
	Service        string             `yaml:"service"`
	ObjectName     string             `yaml:"object_name"`
	Field          string             `yaml:"field"`
	Fields         []SchemaNodeConfig `yaml:"fields"`
	ExcludeMethods []string           `yaml:"exclude_methods"`
	FilterMethods  []string           `yaml:"filter_methods"`
}
type Config struct {
	Name          string            `yaml:"name"`
	OutputPath    string            `yaml:"output_path"`
	OutputPackage string            `yaml:"output_package"`
	Queries       *SchemaNodeConfig `yaml:"queries"`
	Mutations     *SchemaNodeConfig `yaml:"mutations"`
}

type Service struct {
	Name   string
	Fields []string
	Pkg    string
}

type Generator struct {
	services []Service
}

func (g Generator) AddServices(s ...Service) {
	g.services = append(g.services, s...)
}
func (g Generator) headTemplateFuncs() template.FuncMap {
	return nil
}
func (g Generator) headTemplateContext() map[string]interface{} {
	return nil
}
func (g Generator) bodyTemplateFuncs() template.FuncMap {
	return nil
}
func (g Generator) bodyTemplateContext() map[string]interface{} {
	return nil
}
func (g Generator) generateHead() ([]byte, error) {
	buf := new(bytes.Buffer)
	t, err := templatesHeadGohtmlBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get templates/body.gohtml")
	}
	bodyTpl, err := template.New("head").Funcs(g.headTemplateFuncs()).Parse(string(t))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}
	err = bodyTpl.Execute(buf, g.headTemplateContext())
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}
	return buf.Bytes(), nil
}
func (g Generator) generateBody() ([]byte, error) {
	buf := new(bytes.Buffer)
	t, err := templatesBodyGohtmlBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get templates/body.gohtml")
	}
	bodyTpl, err := template.New("head").Funcs(g.headTemplateFuncs()).Parse(string(t))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}
	err = bodyTpl.Execute(buf, g.headTemplateContext())
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}
	return buf.Bytes(), nil
}
func (g Generator) Generate(config Config) error {
	file, err := os.OpenFile(config.OutputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrapf(err, "failed to open file '%s'", config.OutputPath)
	}
	defer file.Close()
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
	// TODO: fix this
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		r = res
	}
	_, err = file.Write(r)
	if err != nil {
		return errors.Wrap(err, "failed to write schema to file")
	}
	return nil
}
