package swagger2gql

import (
	"github.com/pkg/errors"
)

type FieldsConfig struct {
	ContextKey string `mapstructure:"context_key"`
}
type ObjectConfig struct {
	ErrorField string                  `mapstructure:"error_field"`
	Fields     map[string]FieldsConfig `mapstructure:"fields"`
}
type MethodConfig struct {
	Alias       string `mapstructure:"alias"`
	RequestType string `mapstructure:"request_type"` // QUERY | MUTATION
}
type TagConfig struct {
	ClientGoPackage string                  `mapstructure:"client_go_package"`
	Alias           string                  `mapstructure:"alias"`
	Methods         map[string]MethodConfig `mapstructure:"methods"`
}
type Config struct {
	Files []*SwaggerFileConfig `mapstructure:"files"`

	// Global configs for proto files
	Paths          []string                  `mapstructure:"paths"`
	ImportsAliases []map[string]string       `mapstructure:"imports_aliases"`
	Messages       []map[string]ObjectConfig `mapstructure:"messages"`
}
type SwaggerFileConfig struct {
	Name string `mapstructure:"name"`

	Path string `mapstructure:"path"`

	ModelsGoPath string `mapstructure:"models_go_path"`

	OutputPkg  string `mapstructure:"output_package"`
	OutputPath string `mapstructure:"output_path"`

	TagsClientsGoPackages map[string]string `mapstructure:"swagger_go_package"` // go package of protoc generated code

	GQLEnumsPrefix   string `mapstructure:"gql_enums_prefix"`
	GQLMessagePrefix string `mapstructure:"gql_messages_prefix"`

	Tags    map[string]TagConfig      `mapstructure:"tags"`
	Objects []map[string]ObjectConfig `mapstructure:"messages"`
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
func (pc *SwaggerFileConfig) GetTags() map[string]TagConfig {
	if pc == nil {
		return map[string]TagConfig{}
	}
	return pc.Tags
}
func (pc *SwaggerFileConfig) GetObjects() []map[string]ObjectConfig {
	if pc == nil {
		return []map[string]ObjectConfig{}
	}
	return pc.Objects
}
