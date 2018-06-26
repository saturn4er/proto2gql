package proto

import (
	"regexp"

	"github.com/pkg/errors"
)

type FieldsConfig struct {
	ContextKey string `yaml:"context_key"`
}
type MessageConfig struct {
	ErrorField string                  `yaml:"error_field"`
	Fields     map[string]FieldsConfig `yaml:"fields"`
}
type MethodConfig struct {
	Alias       string `yaml:"alias"`
	RequestType string `yaml:"request_type"` // QUERY | MUTATION
}
type ServiceConfig struct {
	Alias   string                  `yaml:"alias"`
	Methods map[string]MethodConfig `yaml:"methods"`
}
type Config struct {
	Files []*ProtoFileConfig `yaml:"files"`

	// Global configs for proto files
	Paths          []string                   `yaml:"paths"`
	ImportsAliases []map[string]string        `yaml:"imports_aliases"`
	Messages       []map[string]MessageConfig `yaml:"messages"`
}
type ProtoFileConfig struct {
	Name string `yaml:"name"`

	Paths          []string            `yaml:"paths"`
	ImportsAliases []map[string]string `yaml:"imports_aliases"`

	ProtoPath string `yaml:"proto_path"`

	OutputPkg  string `yaml:"output_package"`
	OutputPath string `yaml:"output_path"`

	ProtoGoPackage string `yaml:"proto_go_package"` // go package of protoc generated code

	GQLEnumsPrefix   string `yaml:"gql_enums_prefix"`
	GQLMessagePrefix string `yaml:"gql_messages_prefix"`

	Services map[string]ServiceConfig   `yaml:"services"`
	Messages []map[string]MessageConfig `yaml:"messages"`
}

func (pc *ProtoFileConfig) MessageConfig(msgName string) (MessageConfig, error) {
	if pc == nil {
		return MessageConfig{}, nil
	}
	for _, cfgs := range pc.Messages {
		for msgNameRegex, cfg := range cfgs {
			r, err := regexp.Compile(msgNameRegex)
			if err != nil {
				return MessageConfig{}, errors.Wrapf(err, "failed to compile message name regex '%s'", msgNameRegex)
			}
			if r.MatchString(msgName) {
				return cfg, nil
			}
		}
	}
	return MessageConfig{}, nil
}
func (pc *ProtoFileConfig) GetName() string {
	if pc == nil {
		return ""
	}
	return pc.Name
}
func (pc *ProtoFileConfig) GetPaths() []string {
	if pc == nil {
		return []string{}
	}
	return pc.Paths
}
func (pc *ProtoFileConfig) GetProtoPath() string {
	if pc == nil {
		return ""
	}
	return pc.ProtoPath
}
func (pc *ProtoFileConfig) GetOutputPkg() string {
	if pc == nil {
		return ""
	}
	return pc.OutputPkg
}
func (pc *ProtoFileConfig) GetGoPackage() string {
	if pc == nil {
		return ""
	}
	return pc.ProtoGoPackage
}
func (pc *ProtoFileConfig) GetOutputPath() string {
	if pc == nil {
		return ""
	}
	return pc.OutputPath
}
func (pc *ProtoFileConfig) GetGQLEnumsPrefix() string {
	if pc == nil {
		return ""
	}
	return pc.GQLEnumsPrefix
}
func (pc *ProtoFileConfig) GetGQLMessagePrefix() string {
	if pc == nil {
		return ""
	}
	return pc.GQLMessagePrefix
}
func (pc *ProtoFileConfig) GetImportsAliases() []map[string]string {
	if pc == nil {
		return []map[string]string{}
	}
	return pc.ImportsAliases
}
func (pc *ProtoFileConfig) GetServices() map[string]ServiceConfig {
	if pc == nil {
		return map[string]ServiceConfig{}
	}
	return pc.Services
}
func (pc *ProtoFileConfig) GetMessages() []map[string]MessageConfig {
	if pc == nil {
		return []map[string]MessageConfig{}
	}
	return pc.Messages
}

type SchemaNodeConfig struct {
	Type           string             `yaml:"type"` // "OBJECT|SERVICE"
	Proto          string             `yaml:"proto"`
	Service        string             `yaml:"service"`
	ObjectName     string             `yaml:"object_name"`
	Field          string             `yaml:"field"`
	Fields         []SchemaNodeConfig `yaml:"fields"`
	ExcludeMethods []string           `yaml:"exclude_methods"`
	FilterMethods  []string           `yaml:"filter_methods"`
}
type SchemaConfig struct {
	Name          string            `yaml:"name"`
	OutputPath    string            `yaml:"output_path"`
	OutputPackage string            `yaml:"output_package"`
	Queries       *SchemaNodeConfig `yaml:"queries"`
	Mutations     *SchemaNodeConfig `yaml:"mutations"`
}
type GenerateConfig struct {
	Tracer     bool
	VendorPath string
}
