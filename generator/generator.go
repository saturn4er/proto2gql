package generator

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/generator/proto"
)

type generator struct {
	config *GenerateConfig
}

func (g *generator) generateProtos() error {
	protoGen := proto.Generator{
		VendorPath:      g.config.VendorPath,
		GenerateTracers: g.config.GenerateTraces,
	}
	for _, protoFileConfig := range g.config.Protos.Files {
		err := protoGen.AddSourceByConfig(protoFileConfig)
		if err != nil {
			return errors.Wrap(err, "failed to add source to protos generator")
		}
	}
	err := protoGen.Generate()
	if err != nil {
		return errors.Wrap(err, "failed to generate proto derivative files")
	}
	return nil
}
func (g *generator) generate() error {
	err := g.generateProtos()
	if err != nil {
		return errors.Wrap(err, "failed to generate protos")
	}
	// TODO: generate swagger
	// TODO: generate schemas
	return nil
}

func (g *generator) extendProtoFilesConfigs() error {
	for _, cfg := range g.config.Protos.Files {
		cfg.Paths = append(cfg.GetPaths(), g.config.Protos.Paths...)
		for i, path := range cfg.Paths {
			cfg.Paths[i] = os.ExpandEnv(path)
		}
		cfg.ProtoPath = os.ExpandEnv(cfg.ProtoPath)
		cfg.ImportsAliases = append(cfg.ImportsAliases, g.config.Protos.ImportsAliases...)
		cfg.Messages = append(cfg.Messages, g.config.Protos.Messages...)
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
func (g *generator) normalizeConfigPaths() error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to resolve working directory")
	}
	g.config.Protos.Paths = append(g.config.Protos.Paths, wd)

	if g.config.VendorPath != "" {
		vp, err := filepath.Abs(g.config.VendorPath)
		if err != nil {
			return errors.Wrap(err, "failed to normalize vendor path")
		}
		g.config.VendorPath = vp
	}

	return g.extendProtoFilesConfigs()
}

func Generate(gc *GenerateConfig) error {
	var g = generator{
		config: gc,
	}
	err := g.normalizeConfigPaths()
	if err != nil {
		return errors.Wrap(err, "failed to normalize config paths")
	}
	if err := g.generate(); err != nil {
		return errors.Wrap(err, "failed to generate schema")
	}
	return nil
}
