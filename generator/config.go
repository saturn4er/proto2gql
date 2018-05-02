package generator

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
}
type ImportConfig struct {
	GoPackage        string `yaml:"go_package"`
	GQLEnumsPrefix   string `yaml:"gql_enums_prefix"`
	GQLMessagePrefix string `yaml:"gql_messages_prefix"`
}
type ImportsConfig struct {
	OutputPkg  string                  `yaml:"output_package"`
	OutputPath string                  `yaml:"output_path"`
	Aliases    map[string]string       `yaml:"aliases"`
	Settings   map[string]ImportConfig `yaml:"settings"`
}
type GenerateConfig struct {
	Tracer     bool           `yaml:"generate_tracer"`
	VendorPath string         `yaml:"vendor_path"`
	Imports    ImportsConfig  `yaml:"imports"`
	Paths      []string       `yaml:"paths"`
	Protos     []*ProtoConfig `yaml:"protos"`
	OutputPath string         `yaml:"output_path"`
	OutputPkg  string         `yaml:"output_package"`
}
