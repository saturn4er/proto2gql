package graphql

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator"
)

const (
	PluginName      = "graphql_types"
	PluginConfigKey = "graphql_schemas"
)

type Plugin struct {
	files         map[string]*TypesFile
	schemaConfigs []SchemaConfig
}


func (p *Plugin) Prepare() error {
	return nil
}
func (p *Plugin) Init(config *generator.GenerateConfig, plugins []generator.Plugin) error {
	var cfgs []SchemaConfig
	p.files = make(map[string]*TypesFile)
	err := mapstructure.Decode(config.PluginsConfigs[PluginConfigKey], &cfgs)
	if err != nil {
		return errors.Wrap(err, "failed to decode config")
	}
	p.schemaConfigs = cfgs
	return nil
}
func (p *Plugin) AddTypesFile(outputPath string, file *TypesFile) {
	p.files[outputPath] = file
}
func (p Plugin) Name() string {
	return PluginName
}
func (p *Plugin) generateTypes() error{
	for outputPath, file := range p.files {
		err := os.MkdirAll(filepath.Dir(outputPath), 0666)
		if err != nil {
			return errors.Wrapf(err, "failed to create directories for output types file %s", outputPath)
		}
		out, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			return errors.Wrapf(err, "failed to open file %s for write", outputPath)
		}
		err = generateTypes(file, out)
		if err != nil {
			if cerr := out.Close(); cerr != nil {
				err = errors.Wrap(err, cerr.Error())
			}
			return errors.Wrapf(err, "failed to generate types file %s", outputPath)
		}
		if err = out.Close(); err != nil {
			return errors.Wrapf(err, "failed tp close generated types file %s", outputPath)
		}
	}
	return nil
}
func (p *Plugin) Generate() error {
	err := p.generateTypes()
	if err != nil {
		return errors.Wrap(err, "failed to generate types files")
	}
	return nil
}
