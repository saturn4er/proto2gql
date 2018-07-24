// +build NOT (linux ORd darwin)

package main

import (
	"github.com/saturn4er/proto2gql/generator"
	"github.com/saturn4er/proto2gql/generator/plugins/graphql"
	"github.com/saturn4er/proto2gql/generator/plugins/proto2gql"
	"github.com/saturn4er/proto2gql/generator/plugins/swagger2gql"
	"github.com/urfave/cli"
)

func Plugins(c *cli.Context) []generator.Plugin {
	return []generator.Plugin{
		new(graphql.Plugin),
		new(swagger2gql.Plugin),
		new(proto2gql.Plugin),
	}
}
