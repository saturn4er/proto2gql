package parser

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
)

type Parser struct {
	parsedFiles []*File
}

func (p Parser) ParsedFiles() []*File {
	return p.parsedFiles
}

func (p *Parser) Parse(loc string, r io.Reader) (*File, error) {
	fullSwaggerFile, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}
	schema := new(spec.Swagger)
	err = schema.UnmarshalJSON(fullSwaggerFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal File")
	}

	var res = &File{
		file:     schema,
		BasePath: schema.BasePath,
		Location: loc,
	}
	tags, err := parseFileTags(schema, res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file tags")
	}
	res.Tags = tags
	res.file = nil
	return res, nil
}

func resolveSchemaType(route []string, root *spec.Swagger, schema *spec.Schema) (Type, error) {
	if schema == nil {
		return Scalar{kind: KindNull}, nil
	}
	if schema.Ref.String() != "" {
		var err error
		schema, err = spec.ResolveRef(root, &schema.Ref)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve $ref")
		}
	}
	if len(schema.Type) != 1 {
		return nil, errors.Errorf("schema type doesn't contains exactly one element: %v", schema.Type)
	}
	switch schema.Type[0] {
	case "array":
		itemSchema := schema.Items.Schema
		itemType, err := resolveSchemaType(route, root, itemSchema)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve array items types")
		}
		return Array{
			ElemType: itemType,
		}, nil

	case "object":
		if schema.Title != "" {
			route = []string{schema.Title}
		}
		if schema.AdditionalProperties != nil && schema.AdditionalProperties.Schema != nil {
			elemType, err := resolveSchemaType(route, root, schema.AdditionalProperties.Schema)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve hashmap value type")
			}
			typ := Map{
				Route:    route,
				ElemType: elemType,
			}
			return typ, nil
		}
		typ := Object{
			Route: route,
			Name:  schema.Title,
		}

		requiredFields := map[string]struct{}{}
		for _, requiredField := range schema.Required {
			requiredFields[requiredField] = struct{}{}
		}
		for name, prop := range schema.Properties {
			_, required := requiredFields[name]
			ptyp, err := resolveSchemaType(append(route, name), root, &prop)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to resolve prop '%s' type", name)
			}
			typ.Properties = append(typ.Properties, ObjectProperty{
				Name:        name,
				Description: prop.Description,
				Required:    required,
				Type:        ptyp,
			})
		}
		return typ, nil
	case "number":
		switch schema.Format {
		case "float":
			return Scalar{kind: KindFloat32}, nil
		default:
			return Scalar{kind: KindFloat64}, nil
		}

	case "integer":
		switch schema.Format {
		case "int32":
			return Scalar{kind: KindInt32}, nil
		default:
			return Scalar{kind: KindInt64}, nil
		}
	case "boolean":
		return Scalar{kind: KindBoolean}, nil
	case "string":
		if len(schema.Enum) > 0 {
			var values = make([]string, len(schema.Enum))
			for i, enum := range schema.Enum {
				values[i] = enum.(string)
			}
			return Scalar{kind: KindString}, nil
		} else {
			return Scalar{kind: KindString}, nil
		}
	default:
		return nil, errors.Errorf("type %s is not implemented", schema.Type[0])

	}
}
func parseMethodParams(schema *spec.Swagger, method *spec.Operation) ([]MethodParameter, error) {
	var res []MethodParameter
	for _, parameter := range method.Parameters {
		typ, err := resolveSchemaType([]string{method.ID}, schema, parameter.Schema)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve %s parameter type", parameter.Name)
		}
		pos, ok := parameterPositions[parameter.In]
		if !ok {
			return nil, errors.Errorf("unknown parameter position '%s'", parameter.In)
		}
		res = append(res, MethodParameter{
			Name:        parameter.Name,
			Description: parameter.Description,
			Required:    parameter.Required,
			Type:        typ,
			Position:    pos,
		})
	}
	return res, nil
}
func parseMethodResponses(schema *spec.Swagger, method *spec.Operation) ([]MethodResponse, error) {
	var res []MethodResponse
	for statusCode, response := range method.Responses.StatusCodeResponses {
		typ, err := resolveSchemaType([]string{method.ID}, schema, response.Schema)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve schema type")
		}
		res = append(res, MethodResponse{
			StatusCode:  statusCode,
			Description: response.Description,
			ResultType:  typ,
		})
	}
	return res, nil
}
func parseFileTags(schema *spec.Swagger, file *File) ([]Tag, error) {
	var tagsByName = make(map[string]*Tag)
	for _, tag := range schema.Tags {
		tagsByName[tag.Name] = &Tag{
			Name:        tag.Name,
			Description: tag.Description,
		}
	}
	if schema.Paths != nil {
		for path, pathItems := range schema.Paths.Paths {
			methods := map[string]*spec.Operation{
				"GET":     pathItems.Get,
				"PUT":     pathItems.Put,
				"POST":    pathItems.Post,
				"DELETE":  pathItems.Delete,
				"OPTIONS": pathItems.Options,
				"HEAD":    pathItems.Head,
				"PATCH":   pathItems.Patch,
			}
			for httpMethod, method := range methods {
				if method == nil {
					continue
				}
				methodTags := method.Tags
				if len(method.Tags) == 0 {
					methodTags = []string{"operations"}
				}

				params, err := parseMethodParams(schema, method)
				if err != nil {
					return nil, errors.Wrap(err, "failed to parse method params")
				}
				resps, err := parseMethodResponses(schema, method)
				if err != nil {
					return nil, errors.Wrap(err, "failed to resolve method responses")
				}
				m := Method{
					OperationID: method.ID,
					HTTPMethod:  strings.ToUpper(httpMethod),
					Description: method.Description,
					Path:        path,
					Responses:   resps,
					Parameters:  params,
				}
				for _, tag := range methodTags {
					t, ok := tagsByName[tag]
					if !ok {
						t = &Tag{
							Name: tag,
						}
						tagsByName[tag] = t
					}
					t.Methods = append(t.Methods, m)
				}
			}
		}
	}
	var res []Tag
	for _, tag := range tagsByName {
		res = append(res, *tag)
	}
	return res, nil
}
