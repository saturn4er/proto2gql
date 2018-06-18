package common_generator

import (
	"fmt"
	"os"
	"testing"
	"reflect"
	"github.com/saturn4er/proto2gql/common_generator/test_data"
)

type A int

func TestGenerate(t *testing.T) {
	file, err := os.OpenFile("./res/a.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println(Generate(&File{
		PackageName: "hello",
		PackagePath: "github.com/saturn4er/proto2gql/common_generator/res",
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
		OutputObjects: []OutputObject{
			{
				GraphQLName:  "helloOutput",
				VariableName: "test",
				Fields: []ObjectField{
					{
						Name: "hello",
						Type: func(ctx BodyContext) string {
							return ctx.Importer.New(GraphqlPkgPath) + ".String"
						},
						GoObjectType:  reflect.TypeOf(test_data.Hello{}),
						GoObjectField: "Save",
						NeedCast:      true,
						CastTo:        reflect.TypeOf(int(0)),
					},
				},
			},
		},
		MapInputObjects: []MapInputObject{
			{
				VariableName: "someMapInput",
				GraphQLName:  "someMapInput",
				KeyObjectType: func(ctx BodyContext) string {
					return ctx.Importer.New(GraphqlPkgPath) + ".String"
				},
				ValueObjectType: func(ctx BodyContext) string {
					return ctx.Importer.New(GraphqlPkgPath) + ".String"
				},
			},
		},
		MapInputObjectResolvers: []MapInputObjectResolver{
			{
				FunctionName: "ResolveSomeMap",
				KeyGoType:    reflect.TypeOf("1"),
				ValueGoType:  reflect.TypeOf(int32(1)),
				KeyResolver: func(arg string, ctx BodyContext) string {
					return `"hello"`
				},
				ValueResolver: func(arg string, ctx BodyContext) string {
					return `int32(5)`
				},
			},
		},
		MapOutputObjects: []MapOutputObject{
			{
				VariableName: "someMap",
				GraphQLName:  "someMap",
				KeyObjectType: func(ctx BodyContext) string {
					return ctx.Importer.New(GraphqlPkgPath) + ".String"
				},
				KeyGoType: reflect.TypeOf(int(1)),
				ValueObjectType: func(ctx BodyContext) string {
					return ctx.Importer.New(GraphqlPkgPath) + ".String"
				},
				ValueGoType: reflect.TypeOf(int(1)),
			},
		},
		InputObjectResolvers: []InputObjectResolver{
			{
				FunctionName: "HelloResolver",
				OutputGoType: reflect.TypeOf(test_data.Hello{}),
				Fields: []InputObjectResolverField{
					{
						Name:              "AArray",
						GraphqlInputField: "hello",
						GoType:            reflect.TypeOf([]test_data.A{1}),
						ValueResolver: func(arg string, ctx BodyContext) string {
							return ctx.Importer.New("github.com/saturn4er/proto2gql/common_generator/test_data") + ".A(" + arg + ".(int))"
						},
					},
					{
						Name:              "Save",
						GraphqlInputField: "hello2",
						GoType:            reflect.TypeOf(test_data.A(1)),
						ValueResolver: func(arg string, ctx BodyContext) string {
							return ctx.Importer.New("github.com/saturn4er/proto2gql/common_generator/test_data") + ".A(" + arg + ".(int))"
						},
					},
				},
			},
		},
	}, file))
}
