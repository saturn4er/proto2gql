package proto

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
type ProtoConfig struct {
	Name             string                   `yaml:"name"`
	Paths            []string                 `yaml:"paths"`
	ProtoPath        string                   `yaml:"proto_path"`
	OutputPkg        string                   `yaml:"output_package"`
	OutputPath       string                   `yaml:"output_path"`
	GoPackage        string                   `yaml:"go_package"`
	GQLEnumsPrefix   string                   `yaml:"gql_enums_prefix"`
	GQLMessagePrefix string                   `yaml:"gql_messages_prefix"`
	ImportsAliases   map[string]string        `yaml:"imports_aliases"`
	Services         map[string]ServiceConfig `yaml:"services"`
	Messages         map[string]MessageConfig `yaml:"messages"`
}

func (pc *ProtoConfig) GetName() string {
	if pc == nil {
		return ""
	}
	return pc.Name
}
func (pc *ProtoConfig) GetPaths() []string {
	if pc == nil {
		return []string{}
	}
	return pc.Paths
}
func (pc *ProtoConfig) GetProtoPath() string {
	if pc == nil {
		return ""
	}
	return pc.ProtoPath
}
func (pc *ProtoConfig) GetOutputPkg() string {
	if pc == nil {
		return ""
	}
	return pc.OutputPkg
}
func (pc *ProtoConfig) GetGoPackage() string {
	if pc == nil {
		return ""
	}
	return pc.GoPackage
}
func (pc *ProtoConfig) GetOutputPath() string {
	if pc == nil {
		return ""
	}
	return pc.OutputPath
}
func (pc *ProtoConfig) GetGQLEnumsPrefix() string {
	if pc == nil {
		return ""
	}
	return pc.GQLEnumsPrefix
}
func (pc *ProtoConfig) GetGQLMessagePrefix() string {
	if pc == nil {
		return ""
	}
	return pc.GQLMessagePrefix
}
func (pc *ProtoConfig) GetImportsAliases() map[string]string {
	if pc == nil {
		return map[string]string{}
	}
	return pc.ImportsAliases
}
func (pc *ProtoConfig) GetServices() map[string]ServiceConfig {
	if pc == nil {
		return map[string]ServiceConfig{}
	}
	return pc.Services
}
func (pc *ProtoConfig) GetMessages() map[string]MessageConfig {
	if pc == nil {
		return map[string]MessageConfig{}
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
