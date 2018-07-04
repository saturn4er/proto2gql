package generator

import (
	"fmt"

	"github.com/pkg/errors"
)

type Plugin interface {
	Init(*GenerateConfig, []Plugin) error
	Prepare() error
	Name() string
	Generate() error
}
type Generator struct {
	Config  *GenerateConfig
	Plugins []Plugin
}

type PluginContext struct {
	Plugins []Plugin
}

func (g *Generator) RegisterPlugin(p Plugin) error {
	for _, plugin := range g.Plugins {
		if plugin.Name() == p.Name() {
			return errors.Errorf("plugin with name '%s' already exists", p.Name())
		}
	}
	g.Plugins = append(g.Plugins, p)
	return nil
}

func (g *Generator) Generate() error {
	fmt.Println(g.Config.PluginsConfigs)
	for _, plugin := range g.Plugins {
		err := plugin.Init(g.Config, g.Plugins)
		if err != nil {
			return errors.Wrapf(err, "failed to initialize plugin %s", plugin.Name())
		}
	}
	for _, plugin := range g.Plugins {
		err := plugin.Prepare()
		if err != nil {
			return errors.Wrapf(err, "plugin %s preparation error", plugin.Name())
		}
	}
	for _, plugin := range g.Plugins {
		err := plugin.Generate()
		if err != nil {
			return errors.Wrapf(err, "plugin %s generation errors", plugin.Name())
		}
	}
	return nil
}

func Generate(gc *GenerateConfig) error {
	var g = Generator{
		Config: gc,
	}
	if err := g.Generate(); err != nil {
		return errors.Wrap(err, "failed to generate schema")
	}
	return nil
}
