package generator

import (
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v1"
)

const (
	MethodTypeMutation = "MUTATION"
	MethodTypeQuery    = "QUERY"
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
type EnumConfig struct {
	Alias         string            `yaml:"alias"`
	ValuesAliases map[string]string `yaml:"values_aliases"`
}
type ProtoConfig struct {
	Paths            []string                 `yaml:"paths"`
	ProtoPath        string                   `yaml:"proto_path"`
	OutputPkg        string                   `yaml:"output_package"`
	OutputPath       string                   `yaml:"output_path"`
	GQLEnumsPrefix   string                   `yaml:"gql_enums_prefix"`
	GQLMessagePrefix string                   `yaml:"gql_messages_prefix"`
	ImportsAliases   map[string]string        `yaml:"imports_aliases"`
	Services         map[string]ServiceConfig `yaml:"services"`
	Messages         map[string]MessageConfig `yaml:"messages"`
	Enums            map[string]EnumConfig    `yaml:"enums"`
}

func (p ProtoConfig) Copy() ProtoConfig {
	yml, err := yaml.Marshal(p)
	if err != nil {
		panic("Can't marshal ProtoConfig:" + err.Error())
	}
	var res ProtoConfig
	err = yaml.Unmarshal(yml, &res)
	if err != nil {
		panic("Can't unmarshal ProtoConfig:" + err.Error())
	}
	return res
}
func (p *ProtoConfig) Merge(cfg ProtoConfig) {
	err := mergo.Merge(p, cfg)
	panic(err)
}

type ImportsConfig struct {
	OutputPkg  string            `yaml:"output_package"`
	OutputPath string            `yaml:"output_path"`
	Aliases    map[string]string `yaml:"aliases"`
	Settings   map[string]struct {
		GoPackage        string                   `yaml:"go_package"`
		GQLEnumsPrefix   string                   `yaml:"gql_enums_prefix"`
		GQLMessagePrefix string                   `yaml:"gql_messages_prefix"`
		Services         map[string]ServiceConfig `yaml:"services"`
		Enums            map[string]EnumConfig    `yaml:"enums"`
	} `yaml:"settings"`
}
type GenerateConfig struct {
	Tracer     bool           `yaml:"generate_tracer"`
	Imports    ImportsConfig  `yaml:"imports"`
	Paths      []string       `yaml:"paths"`
	Protos     []*ProtoConfig `yaml:"protos"`
	OutputPath string         `yaml:"output_path"`
	OutputPkg  string         `yaml:"output_package"`
}
