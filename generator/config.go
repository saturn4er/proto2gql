package generator

const (
	MethodTypeMutation = "MUTATION"
	MethodTypeQuery    = "QUERY"

	SchemaNodeTypeObject  = "OBJECT"
	SchemaNodeTypeService = "SERVICE"
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
	Name             string                   `yaml:"name"`
	Paths            []string                 `yaml:"paths"`
	ProtoPath        string                   `yaml:"proto_path"`
	OutputPkg        string                   `yaml:"output_package"`
	GoPackage        string                   `yaml:"go_package"`
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
	Tracer     bool           `yaml:"generate_tracer"`
	VendorPath string         `yaml:"vendor_path"`
	Imports    ImportsConfig  `yaml:"imports"`
	Paths      []string       `yaml:"paths"`
	Protos     []*ProtoConfig `yaml:"protos"`
	Schemas    []SchemaConfig `yaml:"schemas"`
	OutputPath string         `yaml:"output_path"`
	OutputPkg  string         `yaml:"output_package"`
}
