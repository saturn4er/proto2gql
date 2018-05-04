package generator

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/parser"
)

type objectObjectField struct {
	Object *gqlObject
}
type fieldConfig struct {
	ServiceFile   *gqlProtoDerivativeFile
	Service       *parser.Service
	QuotedComment string
	Name          string
}
type gqlObject struct {
	QueryObject   bool
	Name          string
	QuotedComment string
	Fields        []fieldConfig
}

type schemaGenerator struct {
	cfg            SchemaConfig
	protos         map[*ProtoConfig]*gqlProtoDerivativeFile
	objects        []*gqlObject
	queryObject    *gqlObject
	mutationObject *gqlObject
}

func (g *schemaGenerator) resolveObjectFields(nodeCfg SchemaNodeConfig, object *gqlObject) (haveFields bool, err error) {
	switch nodeCfg.Type {
	case SchemaNodeTypeObject:
		for _, fld := range nodeCfg.Fields {
			fldObj := &gqlObject{
				QueryObject: object.QueryObject,
				Name:        fld.ObjectName,
			}
			haveFields, err := g.resolveObjectFields(fld, fldObj)
			if err != nil {
				return false, errors.Wrapf(err, "can't resolve field %s object fields", fld.Field)
			}
			if haveFields {
				object.Fields = append(object.Fields, &fieldConfig{
					Name: fld.Field,
				})
			}
		}
		return len(object.Fields) > 0, nil
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
				return false, errors.Errorf("can't find service '%s' in proto file '%s'", nodeCfg.Service, nodeCfg.Proto)
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
			if len(nodeCfg.FilterMethods) > 0 {
				for _, m := range methods {
					var found bool

				}
			}

			object.Fields = objectFields{
				Service: &objectServiceFields{
					File:    file,
					Service: service,
				},
			}
			return true, nil
		}
		return false, errors.Errorf("service '%s' not found", nodeCfg.Service)

	default:
		return false, errors.Errorf("unknown type %s", nodeCfg.Type)
	}
	return false, nil
}
func (g *schemaGenerator) resolveObjectsToGenerate() error {
	var queryObj = &gqlObject{
		QueryObject: true,
		Name:        "Query",
	}
	ok, err := g.resolveObjectFields(g.cfg.Queries, queryObj)
	if err != nil {
		return errors.Wrap(err, "failed to resolve queries fields")
	}

	if ok {
		g.objects = append(g.objects, queryObj)
		g.queryObject = queryObj
	}
	var mutationObj = &gqlObject{
		QueryObject: false,
		Name:        "Mutation",
	}
	ok, err = g.resolveObjectFields(g.cfg.Mutations, mutationObj)
	fmt.Println(mutationObj, err)

	if err != nil {
		return errors.Wrap(err, "failed to resolve mutations fields")
	}
	if ok {
		g.objects = append(g.objects, mutationObj)
		g.mutationObject = mutationObj
	}
	return nil
}
func (g *schemaGenerator) generate() error {
	err := g.resolveObjectsToGenerate()
	if err != nil {
		return errors.Wrap(err, "failed to resolve objects, that we need to generate")
	}

	return nil
}
func generateSchema(sc SchemaConfig, protos map[*ProtoConfig]*gqlProtoDerivativeFile) error {
	g := &schemaGenerator{
		cfg:    sc,
		protos: protos,
	}
	return g.generate()
}
