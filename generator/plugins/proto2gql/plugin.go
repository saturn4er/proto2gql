package proto2gql

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
)

const (
	PluginName      = "proto2gql"
	PluginConfigKey = "proto2gql"
)

type Plugin struct {
	graphql        *graphql.Plugin
	config         *Config
	generateConfig *generator.GenerateConfig
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
	err = p.normalizeGenerateConfigPaths()
	if err != nil {
		return errors.Wrap(err, "failed to normalize config paths")
	}
	return nil
}
func (p *Plugin) normalizeGenerateConfigPaths() error {
	for i, path := range p.config.Paths {
		normalizedPath := os.ExpandEnv(path)
		normalizedPath, err := filepath.Abs(normalizedPath)
		if err != nil {
			return errors.Wrapf(err, "failed to make normalized path '%s' absolute", normalizedPath)
		}
		p.config.Paths[i] = normalizedPath
	}
	return nil
}
func (p *Plugin) prepareFileConfig(fileCfg *ProtoFileConfig) error {
	fileCfg.Paths = append(fileCfg.Paths, p.config.Paths...)
	for _, aliases := range p.config.ImportsAliases {
		fileCfg.ImportsAliases = append(fileCfg.ImportsAliases, aliases)
	}
	return nil
}
func (p *Plugin) Prepare() error {
	pr := new(Proto2GraphQL)
	pr.VendorPath = p.generateConfig.VendorPath
	pr.GenerateTracers = p.generateConfig.GenerateTraces
	for _, file := range p.config.Files {
		err := p.prepareFileConfig(file)
		if err != nil {
			return errors.Wrapf(err, "failed to prepare file %s config", file.ProtoPath)
		}
		err = pr.AddSourceByConfig(file)
		if err != nil {
			return errors.Wrap(err, "failed to parse file "+file.ProtoPath)
		}
	}
	for _, file := range pr.parser.ParsedFiles() {
		pf, err := pr.parsedFile(file)
		if err != nil {
			return errors.Wrapf(err, "failed to resovle parsed file of '%s'", file.FilePath)
		}

		commonFile, err := pr.prepareFile(pf)
		if err != nil {
			return errors.Wrap(err, "failed to prepare file for generation")
		}
		p.graphql.AddTypesFile(pf.OutputPath, commonFile)
	}
	return nil
}

func (g *Proto2GraphQL) AddSourceByConfig(config *ProtoFileConfig) error {
	file, err := g.parser.Parse(config.ProtoPath, config.ImportsAliases, config.Paths)
	if err != nil {
		return errors.Wrap(err, "failed to parse proto file")
	}
	outPath, err := g.fileOutputPath(config, file)
	if err != nil {
		return errors.Wrapf(err, "failed to resolve file '%s' output path", file.FilePath)
	}
	outPkgName, outPkg, err := g.fileOutputPackage(config, file)
	if err != nil {
		return errors.Wrapf(err, "failed to resolve file '%s' output Go package", file.FilePath)
	}
	grpcPkg, err := g.fileGRPCSourcesPackage(config, file)
	if err != nil {
		return errors.Wrapf(err, "failed to resolve file '%s' GRPC sources Go package", file.FilePath)
	}
	g.ParsedFiles = append(g.ParsedFiles, &parsedFile{
		File:           file,
		Config:         config,
		OutputPath:     outPath,
		OutputPkg:      outPkg,
		OutputPkgName:  outPkgName,
		GRPCSourcesPkg: grpcPkg,
	})
	return nil
}

func (Plugin) Name() string {
	return PluginName
}

func (Plugin) Generate() error {
	return nil
}