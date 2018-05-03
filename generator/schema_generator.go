package generator

import "github.com/pkg/errors"

type objectField struct {
	Name          string
	QuotedComment string
	Type          *gqlObject
}
type objectServiceField struct {
	File        *generatedFile
	ServiceName string
}

type objectFields struct {
	Service *objectServiceField
	Fields  []objectField
}
type gqlObject struct {
	Name          string
	QuotedComment string
	Fields        objectFields
}

type schemaGenerator struct {
	cfg            SchemaConfig
	protos         map[*ProtoConfig]*generatedFile
	objects        []*gqlObject
	queryObject    *gqlObject
	mutationObject *gqlObject
}

func (g *schemaGenerator) resolveObjectFields(nc SchemaNodeConfig, object *gqlObject) (haveFields bool, err error) {
	switch nc.Type {
	case SchemaNodeTypeObject:

	case SchemaNodeTypeService:
		for cfg, file := range g.protos {
			if cfg.Name == nc.ProtoName {
				object.Fields = objectFields{
					Service: &objectServiceField{
						File:        file,
						ServiceName: cfg.Name,
					},
				}
				return true, nil
			}
		}
		return false, errors.Errorf("service '%s' not found", nc.ProtoName)

	default:
		return false, errors.Errorf("unknown type %s", nc.Type)
	}
	return false, nil
}
func (g *schemaGenerator) resolveObjectsToGenerate() error {
	var queryObj = &gqlObject{
		Name: "Query",
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
		Name: "Mutation",
	}
	ok, err = g.resolveObjectFields(g.cfg.Mutations, mutationObj)
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
		return nil
	}
	return nil
}
func generateSchema(sc SchemaConfig, protos map[*ProtoConfig]*generatedFile) error {
	g := &schemaGenerator{
		cfg:    sc,
		protos: protos,
	}
	return g.generate()
}
