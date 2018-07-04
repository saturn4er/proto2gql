package graphql_types

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator"
)

const (
	PluginName      = "graphql_types"
	PluginConfigKey = "graphql_schemas"
)

type Plugin struct {
	files []*File
}

func (p *Plugin) Init(config *generator.GenerateConfig, plugins []generator.Plugin) error {
	var cfg []SchemaConfig
	err := mapstructure.Decode(config.PluginsConfigs[PluginConfigKey], &cfg)
	if err != nil {
		return errors.Wrap(err, "failed to decode config")
	}
	return nil
}
func (p *Plugin) AddFile(file *File) error {
	p.files = append(p.files, file)
	return nil
}
func (p Plugin) Name() string {
	return PluginName
}

func (p *Plugin) Generate() error {
	return nil
}
