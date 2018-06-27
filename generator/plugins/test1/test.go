package test1

import (
	"fmt"

	"github.com/saturn4er/proto2gql/generator"
)

const PluginName = "test1"

type Plugin struct {
}

func (Plugin) Init([]generator.Plugin) error {
	fmt.Println("init plugin 1")
	return nil
}

func (Plugin) Name() string {
	return PluginName
}

func (Plugin) Generate() error {
	return nil
}
