package swagger2gql

import (
	"regexp"

	"github.com/pkg/errors"
)

type FieldsConfig struct {
	ContextKey string `mapstructure:"context_key"`
}
type MessageConfig struct {
	ErrorField string                  `mapstructure:"error_field"`
	Fields     map[string]FieldsConfig `mapstructure:"fields"`
}
type MethodConfig struct {
	Alias       string `mapstructure:"alias"`
	RequestType string `mapstructure:"request_type"` // QUERY | MUTATION
}
type ServiceConfig struct {
	Alias   string                  `mapstructure:"alias"`
	Methods map[string]MethodConfig `mapstructure:"methods"`
}
type Config struct {
	Files []*SwaggerFileConfig `mapstructure:"files"`

	// Global configs for proto files
	Paths          []string                   `mapstructure:"paths"`
	ImportsAliases []map[string]string        `mapstructure:"imports_aliases"`
	Messages       []map[string]MessageConfig `mapstructure:"messages"`
}
type SwaggerFileConfig struct {
	Name string `mapstructure:"name"`

	Path string `mapstructure:"path"`

	OutputPkg  string `mapstructure:"output_package"`
	OutputPath string `mapstructure:"output_path"`

	TagsClientsGoPackages map[string]string `mapstructure:"swagger_go_package"` // go package of protoc generated code

	GQLEnumsPrefix   string `mapstructure:"gql_enums_prefix"`
	GQLMessagePrefix string `mapstructure:"gql_messages_prefix"`

	Services map[string]ServiceConfig   `mapstructure:"services"`
	Messages []map[string]MessageConfig `mapstructure:"messages"`
}

func (pc *SwaggerFileConfig) MessageConfig(msgName string) (MessageConfig, error) {
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
func (pc *SwaggerFileConfig) GetName() string {
	if pc == nil {
		return ""
	}
	return pc.Name
}
func (pc *SwaggerFileConfig) GetPath() string {
	if pc == nil {
		return ""
	}
	return pc.Path
}
func (pc *SwaggerFileConfig) GetOutputPkg() string {
	if pc == nil {
		return ""
	}
	return pc.OutputPkg
}
func (pc *SwaggerFileConfig) GetTagGoPackage(tag string) (string, error) {
	if pc == nil {
		return "", errors.Errorf("go package is not specified for tag '%ws'", tag)
	}
	pkg, ok := pc.TagsClientsGoPackages[tag]
	if !ok {
		return "", errors.Errorf("go package is not specified for tag '%ws'", tag)
	}
	return pkg, nil

}
func (pc *SwaggerFileConfig) GetOutputPath() string {
	if pc == nil {
		return ""
	}
	return pc.OutputPath
}
func (pc *SwaggerFileConfig) GetGQLEnumsPrefix() string {
	if pc == nil {
		return ""
	}
	return pc.GQLEnumsPrefix
}
func (pc *SwaggerFileConfig) GetGQLMessagePrefix() string {
	if pc == nil {
		return ""
	}
	return pc.GQLMessagePrefix
}
func (pc *SwaggerFileConfig) GetServices() map[string]ServiceConfig {
	if pc == nil {
		return map[string]ServiceConfig{}
	}
	return pc.Services
}
func (pc *SwaggerFileConfig) GetMessages() []map[string]MessageConfig {
	if pc == nil {
		return []map[string]MessageConfig{}
	}
	return pc.Messages
}

type SchemaNodeConfig struct {
	Type           string             `mapstructure:"type"` // "OBJECT|SERVICE"
	Proto          string             `mapstructure:"proto"`
	Service        string             `mapstructure:"service"`
	ObjectName     string             `mapstructure:"object_name"`
	Field          string             `mapstructure:"field"`
	Fields         []SchemaNodeConfig `mapstructure:"fields"`
	ExcludeMethods []string           `mapstructure:"exclude_methods"`
	FilterMethods  []string           `mapstructure:"filter_methods"`
}
type SchemaConfig struct {
	Name          string            `mapstructure:"name"`
	OutputPath    string            `mapstructure:"output_path"`
	OutputPackage string            `mapstructure:"output_package"`
	Queries       *SchemaNodeConfig `mapstructure:"queries"`
	Mutations     *SchemaNodeConfig `mapstructure:"mutations"`
}
type GenerateConfig struct {
	Tracer     bool
	VendorPath string
}
