package generator

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/proto"
)

type Plugin interface {
	Init([]Plugin) error
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
func (g *Generator) generateProtos() (*proto.Generator, error) {
	protoGen := &proto.Generator{
		VendorPath:      g.Config.VendorPath,
		GenerateTracers: g.Config.GenerateTraces,
	}
	for _, protoFileConfig := range g.Config.Protos.Files {
		err := protoGen.AddSourceByConfig(protoFileConfig)
		if err != nil {
			return nil, errors.Wrap(err, "failed to add source to protos Generator")
		}
	}
	err := protoGen.Generate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate proto derivative files")
	}
	return protoGen, nil
}

func (g *Generator) Generate() error {
	for _, plugin := range g.Plugins {
		err := plugin.Init(g.Plugins)
		if err != nil {
			return errors.Wrapf(err, "failed to initialize plugin %s", plugin.Name())
		}
	}
	for _, plugin := range g.Plugins {
		err := plugin.Generate()
		if err != nil {
			return errors.Wrapf(err, "generation error by plugin %s", plugin.Name())
		}
	}
	return nil
	// err := g.normalizeConfigPaths()
	// if err != nil {
	// 	return errors.Wrap(err, "failed to normalize Config paths")
	// }
	// protos, err := g.generateProtos()
	// if err != nil {
	// 	return errors.Wrap(err, "failed to generate protos")
	// }
	// schemaGen := schema.Generator{}
	// schemaGen.AddServices(protos.SchemaServices()...)
	// for _, schemaConfig := range g.Config.Shemas {
	// 	if err := schemaGen.Generate(schemaConfig); err != nil {
	// 		return errors.Wrapf(err, "failed to generate schema '%s'", schemaConfig.Name)
	// 	}
	// }

	// TODO: generate swagger
	return nil
}

func (g *Generator) extendProtoFilesConfigs() error {
	for _, cfg := range g.Config.Protos.Files {
		cfg.Paths = append(cfg.GetPaths(), g.Config.Protos.Paths...)
		for i, path := range cfg.Paths {
			cfg.Paths[i] = os.ExpandEnv(path)
		}
		cfg.ProtoPath = os.ExpandEnv(cfg.ProtoPath)
		cfg.ImportsAliases = append(cfg.ImportsAliases, g.Config.Protos.ImportsAliases...)
		cfg.Messages = append(cfg.Messages, g.Config.Protos.Messages...)
		if cfg.OutputPath != "" {
			absoluteOutPath, err := filepath.Abs(os.ExpandEnv(cfg.OutputPath))
			if err != nil {
				return errors.Wrap(err, "failed to normalize proto output path")
			}
			cfg.OutputPath = absoluteOutPath
		}

	}
	return nil
}
func (g *Generator) normalizeConfigPaths() error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to resolve working directory")
	}
	g.Config.Protos.Paths = append(g.Config.Protos.Paths, wd)

	if g.Config.VendorPath != "" {
		vp, err := filepath.Abs(g.Config.VendorPath)
		if err != nil {
			return errors.Wrap(err, "failed to normalize vendor path")
		}
		g.Config.VendorPath = vp
	}

	return g.extendProtoFilesConfigs()
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
