package generator

import (
	"github.com/saturn4er/proto2gql/generator/proto"
	"github.com/saturn4er/proto2gql/generator/schema"
)


type GenerateConfig struct {
	GenerateTraces bool            `yaml:"generate_tracer"`
	VendorPath     string          `yaml:"vendor_path"`
	Protos         proto.Config    `yaml:"protos"`
	Shemas         []schema.Config `yaml:"schemas"`
}
