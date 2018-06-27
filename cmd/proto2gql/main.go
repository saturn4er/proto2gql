package main

import (
	"io/ioutil"
	"os"

	"github.com/saturn4er/proto2gql/generator"
	"github.com/saturn4er/proto2gql/generator/plugins/test1"
	"github.com/saturn4er/proto2gql/generator/plugins/test2"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
)

func main() {
	app := cli.App{
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Value: "generate.yml",
			},
		},
		Action: func(c *cli.Context) {
			cfgFile, err := os.Open(c.String("config"))
			if err != nil {
				panic(err)
			}
			cfg, err := ioutil.ReadAll(cfgFile)
			if err != nil {
				panic(err)
			}
			gc := new(generator.GenerateConfig)
			err = yaml.Unmarshal(cfg, gc)
			if err != nil {
				panic(err)
			}
			g := &generator.Generator{
				Config: gc,
			}
			plugins := []generator.Plugin{
				new(test1.Plugin),
				new(test2.Plugin),
			}
			for _, plugin := range plugins {
				err := g.RegisterPlugin(plugin)
				if err != nil {
					panic(err.Error())
				}
			}
			err = g.Generate()
			if err != nil {
				panic(err)
			}
		},
	}
	app.Run(os.Args)
}
