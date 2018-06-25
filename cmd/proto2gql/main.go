package main

import (
	"io/ioutil"
	"os"

	"github.com/saturn4er/proto2gql/generator"
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
			err = generator.Generate(gc)
			if err != nil {
				panic(err)
			}
		},
	}
	app.Run(os.Args)
}
