package test2

import (
	"fmt"

	"github.com/saturn4er/proto2gql/generator"
	"github.com/saturn4er/proto2gql/generator/plugins/test1"
)

const PluginName = "test2"

type Plugin struct {
	common *test1.Plugin
}

func (p *Plugin) Init(plugins []generator.Plugin) error {
	for _, plugin := range plugins {
		switch plugin.Name() {
		case test1.PluginName:
			p.common = plugin.(*test1.Plugin)
		}
	}
	fmt.Println(p.common.Generate())
	return nil
}

func (Plugin) Name() string {
	return PluginName
}

func (Plugin) Generate() error {
	return nil
}
