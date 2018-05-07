package generator

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/parser"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type schemaService struct {
	ServiceFile *gqlProtoDerivativeFile
	Service     *parser.Service
}
type fieldConfig struct {
	Service       *schemaService
	QuotedComment string
	Name          string
	Object        *gqlObject
}
type gqlObject struct {
	QueryObject   bool
	Name          string
	QuotedComment string
	Fields        []fieldConfig
}

type schemaGenerator struct {
	cfg            SchemaConfig
	generatorCfg   *GenerateConfig
	protos         map[*ProtoConfig]*gqlProtoDerivativeFile
	objects        []*gqlObject
	queryObject    *gqlObject
	mutationObject *gqlObject
	services       []*schemaService
	usedServices   map[*parser.Service]struct{}
}

func (g *schemaGenerator) filterMethods(methods []*parser.Method, filter, exclude []string) []*parser.Method {
	var res []*parser.Method
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
			if _, ok := excludedMethods[m.Name]; ok {
				continue
			}
		}
		if len(filteredMethods) > 0 {
			if _, ok := filteredMethods[m.Name]; !ok {
				continue
			}
		}
		res = append(res, m)
	}
	return res
}
func (g *schemaGenerator) resolveObjectFields(nodeCfg SchemaNodeConfig, object *gqlObject) (err error) {
	switch nodeCfg.Type {
	case SchemaNodeTypeObject:
		for _, fld := range nodeCfg.Fields {
			fldObj := &gqlObject{
				QueryObject: object.QueryObject,
				Name:        strings.Replace(fld.ObjectName, " ", "_", -1),
			}
			err := g.resolveObjectFields(fld, fldObj)
			if err != nil {
				return errors.Wrapf(err, "can't resolve field %s object fields", fld.Field)
			}
			if len(fldObj.Fields) > 0 {
				object.Fields = append(object.Fields, fieldConfig{
					Name:   fld.Field,
					Object: fldObj,
				})
				g.objects = append(g.objects, fldObj)
			}
		}
		return nil
	case SchemaNodeTypeService:
		for cfg, file := range g.protos {
			if cfg.Name != nodeCfg.Proto {
				continue
			}
			var service *parser.Service
			for _, s := range file.ProtoFile.Services {
				if s.Name == nodeCfg.Service {
					service = s
					break
				}
			}
			if service == nil {
				return errors.Errorf("can't find service '%s' in proto file '%s'", nodeCfg.Service, nodeCfg.Proto)
			}
			var methods []*parser.Method
			for _, m := range service.Methods {
				if file.Generator.methodIsQuery(m) {
					if object.QueryObject {
						methods = append(methods, m)
					}
				} else {
					if !object.QueryObject {
						methods = append(methods, m)
					}
				}
			}
			methods = g.filterMethods(methods, nodeCfg.FilterMethods, nodeCfg.ExcludeMethods)

			schemaService := &schemaService{Service: service, ServiceFile: file}
			if _, ok := g.usedServices[service]; !ok {
				g.usedServices[service] = struct{}{}
				g.services = append(g.services, schemaService)
			}
			for _, m := range methods {
				object.Fields = append(object.Fields, fieldConfig{
					Name:    file.Generator.methodName(m),
					Service: schemaService,
				})
			}
			return nil
		}
		return errors.Errorf("service '%s' not found", nodeCfg.Service)

	default:
		return errors.Errorf("unknown type %s", nodeCfg.Type)
	}
}
func (g *schemaGenerator) resolveObjectsToGenerate() error {
	if g.cfg.Queries != nil {
		var queryObj = &gqlObject{
			QueryObject: true,
			Name:        "Query",
		}
		err := g.resolveObjectFields(*g.cfg.Queries, queryObj)
		if err != nil {
			return errors.Wrap(err, "failed to resolve queries fields")
		}

		g.objects = append(g.objects, queryObj)
		g.queryObject = queryObj
	}
	if g.cfg.Mutations != nil {
		var mutationObj = &gqlObject{
			QueryObject: false,
			Name:        "Mutation",
		}
		err := g.resolveObjectFields(*g.cfg.Mutations, mutationObj)

		if err != nil {
			return errors.Wrap(err, "failed to resolve mutations fields")
		}
		if len(mutationObj.Fields) > 0 {
			g.objects = append(g.objects, mutationObj)
			g.mutationObject = mutationObj
		}
	}
	return nil
}
func (g *schemaGenerator) serviceGoClientType(imports *importer) func(service *schemaService) string {
	return func(service *schemaService) string {
		return imports.New(service.ServiceFile.GoProtoPkg) + "." + camelCase(service.Service.Name) + "Client"
	}
}
func (g *schemaGenerator) serviceGoMutationsFieldsGetter(imports *importer) func(service *schemaService) string {
	return func(service *schemaService) string {
		return imports.New(service.ServiceFile.OutGoPkg) + ".Get" + service.Service.Name + "GraphQLMutationsFields"
	}
}
func (g *schemaGenerator) serviceGoQueriesFieldsGetter(imports *importer) func(service *schemaService) string {
	return func(service *schemaService) string {
		return imports.New(service.ServiceFile.OutGoPkg) + ".Get" + service.Service.Name + "GraphQLQueriesFields"
	}
}
func (g *schemaGenerator) templateContext(imports *importer) map[string]interface{} {
	return map[string]interface{}{
		"pkg":             g.cfg.OutputPackage,
		"schemaName":      g.cfg.Name,
		"objects":         g.objects,
		"queryObj":        g.queryObject,
		"mutationObj":     g.mutationObject,
		"services":        g.services,
		"trace":           g.generatorCfg.Tracer,
		"gqlpkg":          imports.New(graphqlPkgPath),
		"interceptorspkg": imports.New(interceptorsPkgPath),
		"tracerpkg":       imports.New(tracerPkg),
		"imports":         imports.Imports(),

		"CCase":                          camelCase,
		"ServiceGoClientType":            g.serviceGoClientType(imports),
		"ServiceGoQueriesFieldsGetter":   g.serviceGoQueriesFieldsGetter(imports),
		"ServiceGoMutationsFieldsGetter": g.serviceGoMutationsFieldsGetter(imports),
	}
}
func (g *schemaGenerator) generate() error {
	err := g.resolveObjectsToGenerate()
	if err != nil {
		return errors.Wrap(err, "failed to resolve objects, that we need to generate")
	}
	tpl, err := template.New("template").Parse(schemaBodyTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to parse template")
	}
	var imprts = new(importer)
	res := bytes.NewBuffer(nil)
	err = tpl.Execute(res, g.templateContext(imprts))
	if err != nil {
		return errors.Wrap(err, "failed to execute template")
	}
	headres := bytes.NewBuffer(nil)
	hdtpd, err := template.New("header").Parse(schemaHeadTemplate)
	if err != nil {
		panic(err)
	}
	err = hdtpd.Execute(headres, g.templateContext(imprts))
	if err != nil {
		panic(err)
	}
	r := bytes.Join([][]byte{headres.Bytes(), res.Bytes()}, nil)
	r, err = imports.Process(g.cfg.OutputPath, r, &imports.Options{
		Comments: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to format generated code")
	}
	err = os.MkdirAll(filepath.Dir(g.cfg.OutputPath), 0777)
	if err != nil {
		panic(err)
	}
	return ioutil.WriteFile(g.cfg.OutputPath, r, 0600)
}
func generateSchema(cfg *GenerateConfig, sc SchemaConfig, protos map[*ProtoConfig]*gqlProtoDerivativeFile) error {
	g := &schemaGenerator{
		cfg:          sc,
		generatorCfg: cfg,
		protos:       protos,
		usedServices: make(map[*parser.Service]struct{}),
	}
	return g.generate()
}
