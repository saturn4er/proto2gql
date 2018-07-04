package swagger2gql

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql/parser"
)

const (
	PluginName      = "proto2gql"
	PluginConfigKey = "proto2gql"
)

type Plugin struct {
	graphql        *graphql.Plugin
	config         *Config
	generateConfig *generator.GenerateConfig

	parsedFiles []*parsedFile
}

func (p *Plugin) Init(config *generator.GenerateConfig, plugins []generator.Plugin) error {
	for _, plugin := range plugins {
		switch plugin.Name() {
		case graphql.PluginName:
			p.graphql = plugin.(*graphql.Plugin)
		}
	}
	if p.graphql == nil {
		return errors.New("'graphql' plugin is not installed.")
	}
	cfg := new(Config)
	err := mapstructure.Decode(config.PluginsConfigs[PluginConfigKey], cfg)
	if err != nil {
		return errors.Wrap(err, "failed to decode config")
	}
	p.generateConfig = config
	p.config = cfg
	return nil
}
func (p *Plugin) prepareTypesFile(file *parsedFile) (*graphql.TypesFile, error) {
	enums, err := p.prepareFileEnums(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file enums")
	}
	// inputs, err := p.fileInputObjects(file)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to prepare file input objects")
	// }
	// mapInputs, err := p.fileMapInputObjects(file)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to prepare file map input objects")
	// }
	// mapOutputs, err := p.fileMapOutputObjects(file)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to prepare file map output objects")
	// }
	// mapResolvers, err := p.fileInputMapResolvers(file)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to prepare file map resolvers")
	// }
	// outputMessages, err := p.fileOutputMessages(file)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to prepare file output messages")
	// }
	// messagesResolvers, err := p.fileInputMessagesResolvers(file)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to prepare file messages resolvers")
	// }
	// services, err := p.fileServices(file)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to prepare file services")
	// }
	res := &graphql.TypesFile{
		PackageName: file.OutputPkgName,
		Package:     file.OutputPkg,
		// Enums:                   enums,
		// InputObjects:            inputs,
		// InputObjectResolvers:    messagesResolvers,
		// OutputObjects:           outputMessages,
		// MapInputObjects:         mapInputs,
		// MapInputObjectResolvers: mapResolvers,
		// MapOutputObjects:        mapOutputs,
		// Services:                services,
	}
	return res, nil
}
func (p *Plugin) Prepare() error {
	fmt.Println(p.config.Files)
	for _, cfg := range p.config.Files {
		pf, err := parser.Parse(cfg.Path)
		if err != nil {
			return errors.Wrap(err, "failed to parse swagger config")
		}
		outPath, err := p.fileOutputPath(cfg)
		if err != nil {
			return errors.Wrapf(err, "failed to resolve cfg '%s' output path", cfg.Path)
		}
		outPkgName, outPkg, err := p.fileOutputPackage(cfg)
		if err != nil {
			return errors.Wrapf(err, "failed to resolve cfg '%s' output Go package", cfg.Path)
		}
		f := &parsedFile{
			File:          pf,
			Config:        cfg,
			OutputPath:    outPath,
			OutputPkg:     outPkg,
			OutputPkgName: outPkgName,
		}
		gqlFile, err := p.prepareTypesFile(f)
		if err != nil {
			return errors.Wrap(err, "failed to prepare types cfg")
		}
		p.graphql.AddTypesFile("./hello.go", gqlFile)
	}

	return nil
}

func (Plugin) Name() string {
	return PluginName
}

func (Plugin) Generate() error {
	return nil
}
