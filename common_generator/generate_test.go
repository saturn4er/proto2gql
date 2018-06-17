package common_generator

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	file, err := os.OpenFile("./res/a.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println(Generate(&File{
		Package: "hello",
		Enums: []Enum{
			{
				GraphQLName:  "hello",
				VariableName: "helloEnum",
				Values: []EnumValue{
					{
						Name:    "1",
						Value:   0,
						Comment: `"Hello"`,
					},
					{
						Name:    "2",
						Value:   1,
						Comment: `"Hello 1"`,
					},
				},
			},
		},
		InputObjects: []InputObject{
			{
				GraphQLName:  "hello",
				VariableName: "hh",
				Fields: []ObjectField{
					{
						Name: "hello",
						Type: func(ctx BodyContext) string {
							return ctx.Importer.New(GraphqlPkgPath) + ".String"
						},
					},
				},
			},
		},
		InputObjectResolvers: []InputObjectResolver{
			{
				FunctionName: "HelloResolver",
				OutputGoType: "int",
			},
		},
	}, file))
}
