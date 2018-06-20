package generator

import (
	"path/filepath"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/parser"
	"github.com/saturn4er/proto2gql/generator/proto"
	"os"
)

type gqlProtoDerivativeFile struct {
	GoProtoPkg string

	OutGoPkgName string
	OutGoPkg     string

	OutDir      string
	OutFilePath string

	ProtoFile *parser.File

	TracerEnabled    bool
	GQLEnumsPrefix   string
	GQLMessagePrefix string
	Services         map[string]ServiceConfig
	Messages         map[string]MessageConfig
}

type generator struct {
	config *GenerateConfig
}

func (g *generator) generate() error {
	protoGen := proto.Generator{}
	for _, p := range g.config.Protos {
		err := protoGen.AddSourceByConfig(*p)
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

func (g *generator) normalizeConfigPaths() error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to resolve working directory")
	}
	g.config.Paths = append(g.config.Paths, wd)

	if g.config.VendorPath != "" {
		vp, err := filepath.Abs(g.config.VendorPath)
		if err != nil {
			return errors.Wrap(err, "failed to normalize vendor path")
		}
		g.config.VendorPath = vp
	}
	for _, cfg := range g.config.Protos {
		cfg.Paths = append(cfg.Paths, g.config.Paths...)
		if cfg.OutputPath != "" {
			vp, err := filepath.Abs(cfg.OutputPath)
			if err != nil {
				return errors.Wrap(err, "failed to normalize proto output path")
			}
			cfg.OutputPath = vp
		}
	}
	return nil
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
