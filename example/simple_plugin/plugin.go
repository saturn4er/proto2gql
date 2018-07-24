package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
)

const Name = "gql_services_rules"

type plugin struct {
	gqlPlugin *graphql.Plugin
}

func (p *plugin) Init(_ *generator.GenerateConfig, plugins []generator.Plugin) error {
	for _, plugin := range plugins {
		if g, ok := plugin.(*graphql.Plugin); ok {
			p.gqlPlugin = g
			return nil
		}
	}
	return errors.New("graphql plugin was not found")
}

func (p plugin) Generate() error {
	file, err := os.OpenFile("./services_access.yml", os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "failed to open services_access.yml")
	}
	defer file.Close()
	for _, typesFile := range p.gqlPlugin.Types() {
		for _, service := range typesFile.Services {
			if len(service.Methods) == 0 {
				continue
			}
			file.WriteString(service.Name + ":\n")
			for _, method := range service.Methods {
				file.WriteString("   " + method.Name + ":\n")
			}
		}
	}
	return nil
}
func (plugin) Name() string                   { return Name }
func (plugin) Prepare() error                 { return nil }
func (plugin) PrintInfo(info generator.Infos) {}
func (plugin) Infos() map[string]string       { return nil }

func Plugin() generator.Plugin {
	return new(plugin)
}
