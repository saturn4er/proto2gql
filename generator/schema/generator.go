package schema

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/importer"
	"golang.org/x/tools/imports"
)

const (
	GraphqlPkgPath      = "github.com/graphql-go/graphql"
	OpentracingPkgPath  = "github.com/opentracing/opentracing-go"
	TracerPkgPath       = "github.com/saturn4er/proto2gql/api/tracer"
	InterceptorsPkgPath = "github.com/saturn4er/proto2gql/api/interceptors"
	ErrorsPkgPath       = "github.com/pkg/errors"

	SchemaNodeTypeObject  = "OBJECT"
	SchemaNodeTypeService = "SERVICE"
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
	Name            string
	ConstructorName string
	TracerEnabled   bool
	Fields          []string
	Pkg             string
	ClientGoType    GoType
}

type fieldConfig struct {
	QuotedComment string
	Name          string
	Service       *Service
	Object        *gqlObject
}
type gqlObject struct {
	QueryObject   bool
	Name          string
	QuotedComment string
	Fields        []fieldConfig
}

type Generator struct {
	services []Service
}

func (g *Generator) AddServices(s ...Service) {
	g.services = append(g.services, s...)
}
func (g *Generator) importFunc(imports *importer.Importer, importPath string) func() string {
	return func() string {
		return imports.New(importPath)
	}
}
func (g *Generator) resolveObjectFields(nodeCfg SchemaNodeConfig, object *gqlObject) (newObjectx []*gqlObject, err error) {
	var newObjs []*gqlObject
	switch nodeCfg.Type {
	case SchemaNodeTypeObject:
		for _, fld := range nodeCfg.Fields {
			fldObj := &gqlObject{
				QueryObject: object.QueryObject,
				Name:        strings.Replace(fld.ObjectName, " ", "_", -1),
			}
			subObjs, err := g.resolveObjectFields(fld, fldObj)
			if err != nil {
				return nil, errors.Wrapf(err, "can't resolve field %s object fields", fld.Field)
			}
			if len(fldObj.Fields) > 0 {
				object.Fields = append(object.Fields, fieldConfig{
					Name:   fld.Field,
					Object: fldObj,
				})
				newObjs = append(newObjs, subObjs...)
				newObjs = append(newObjs, fldObj)
			}
		}
		return newObjs, nil
	case SchemaNodeTypeService:
		for _, service := range g.services {
			if service.Name != nodeCfg.Service {
				continue
			}
			fields := g.filterMethods(service.Fields, nodeCfg.FilterMethods, nodeCfg.ExcludeMethods)
			for _, fld := range fields {
				object.Fields = append(object.Fields, fieldConfig{
					Name:    fld,
					Service: &service,
				})
			}
			return nil, nil
		}
		return nil, errors.Errorf("service '%s' not found", nodeCfg.Service)

	default:
		return nil, errors.Errorf("unknown type %s", nodeCfg.Type)
	}
}
func (g *Generator) filterMethods(methods []string, filter, exclude []string) []string {
	var res []string
	var filteredMethods = make(map[string]struct{})
	for _, f := range filter {
		filteredMethods[f] = struct{}{}
	}
	var excludedMethods = make(map[string]interface{})
	for _, f := range exclude {
		excludedMethods[f] = struct{}{}
	}
	for _, m := range methods {
		if len(excludedMethods) > 0 {
			if _, ok := excludedMethods[m]; ok {
				continue
			}
		}
		if len(filteredMethods) > 0 {
			if _, ok := filteredMethods[m]; !ok {
				continue
			}
		}
		res = append(res, m)
	}
	return res
}
func (g *Generator) resolveObjectsToGenerate(cfg Config) ([]*gqlObject, string, string, error) {
	var objects []*gqlObject
	var queryObject, mutationObject string
	if cfg.Queries != nil {
		var queryObj = &gqlObject{
			QueryObject: true,
			Name:        "Query",
		}
		newObjs, err := g.resolveObjectFields(*cfg.Queries, queryObj)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "failed to resolve queries fields")
		}
		objects = append(objects, newObjs...)
		objects = append(objects, queryObj)
		queryObject = queryObj.Name
	}
	if cfg.Mutations != nil {
		var mutationObj = &gqlObject{
			QueryObject: false,
			Name:        "Mutation",
		}
		newObjs, err := g.resolveObjectFields(*cfg.Mutations, mutationObj)

		if err != nil {
			return nil, "", "", errors.Wrap(err, "failed to resolve mutations fields")
		}
		if len(mutationObj.Fields) > 0 {
			objects = append(objects, newObjs...)
			objects = append(objects, mutationObj)
			mutationObject = mutationObj.Name
		}
	}
	return objects, queryObject, mutationObject, nil
}
func (g *Generator) headTemplateFuncs(imports *importer.Importer) template.FuncMap {
	return nil
}
func (g *Generator) headTemplateContext(cfg Config, imports *importer.Importer) map[string]interface{} {
	return map[string]interface{}{
		"package": cfg.OutputPackage,
		"imports": imports.Imports(),
	}
}
func (g *Generator) bodyTemplateFuncs(imports *importer.Importer) template.FuncMap {
	return map[string]interface{}{
		"gqlPkg":          g.importFunc(imports, GraphqlPkgPath),
		"tracerPkg":       g.importFunc(imports, TracerPkgPath),
		"interceptorsPkg": g.importFunc(imports, InterceptorsPkgPath),
		"serviceConstructor": func(s *Service) string {
			return imports.Prefix(s.Pkg) + "Get" + s.Name + "ServiceMethods"
		},
		"goType": func(typ GoType) string {
			return typ.String(imports)
		},
	}
}
func (g *Generator) bodyTemplateContext(cfg Config) (map[string]interface{}, error) {
	objects, queryObject, mutationsObj, err := g.resolveObjectsToGenerate(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve objects to generate")
	}
	return map[string]interface{}{
		"schemaName":     cfg.Name,
		"services":       g.services,
		"objects":        objects,
		"queryObject":    queryObject,
		"mutationObject": mutationsObj,
	}, nil
}
func (g *Generator) generateHead(cfg Config, imports *importer.Importer) ([]byte, error) {
	buf := new(bytes.Buffer)
	t, err := templatesHeadGohtmlBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get templates/body.gohtml")
	}
	bodyTpl, err := template.New("head").Funcs(g.headTemplateFuncs(imports)).Parse(string(t))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}
	err = bodyTpl.Execute(buf, g.headTemplateContext(cfg, imports))
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}
	return buf.Bytes(), nil
}
func (g *Generator) generateBody(cfg Config, imports *importer.Importer) ([]byte, error) {
	buf := new(bytes.Buffer)
	t, err := templatesBodyGohtmlBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get templates/body.gohtml")
	}
	bodyTpl, err := template.New("head").Funcs(g.bodyTemplateFuncs(imports)).Parse(string(t))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}
	ctx, err := g.bodyTemplateContext(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare body template context")
	}
	err = bodyTpl.Execute(buf, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}
	return buf.Bytes(), nil
}
func (g *Generator) Generate(config Config) error {
	imprt := new(importer.Importer)
	file, err := os.OpenFile(config.OutputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrapf(err, "failed to open file '%s'", config.OutputPath)
	}
	defer file.Close()
	body, err := g.generateBody(config, imprt)
	if err != nil {
		return errors.Wrap(err, "failed to generate body")
	}
	head, err := g.generateHead(config, imprt)
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
		_, err = file.Write(r)
		if err != nil {
			return errors.Wrap(err, "failed to write non-formatted schema to file")
		}
		return err
	} else {
		r = res
	}
	_, err = file.Write(r)
	if err != nil {
		return errors.Wrap(err, "failed to write formatted schema to file")
	}
	return nil
}
