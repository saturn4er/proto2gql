// This file was generated by github.com/saturn4er/proto2gql. DO NOT EDIT IT

package timestamp

import (
	context "context"
	errors "errors"
	fmt "fmt"
	debug "runtime/debug"
	strconv "strconv"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	graphql "github.com/graphql-go/graphql"
	opentracing_go "github.com/opentracing/opentracing-go"
	interceptors "github.com/saturn4er/proto2gql/api/interceptors"
	scalars "github.com/saturn4er/proto2gql/api/scalars"
)

var (
	_ = errors.New
	_ = graphql.NewObject
	_ = context.Background
	_ = strconv.FormatBool
	_ = fmt.Print
	_ = opentracing_go.GlobalTracer
	_ = debug.FreeOSMemory
)

type (
	_ = interceptors.CallInterceptor
)

// Enums

// Messages
var Timestamp = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Timestamp",
	Fields: graphql.Fields{},
})
var TimestampInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TimestampInput",
	Fields: graphql.InputObjectConfigFieldMapThunk(func() graphql.InputObjectConfigFieldMap {
		return graphql.InputObjectConfigFieldMap{
			"seconds": &graphql.InputObjectFieldConfig{
				Type: scalars.GraphQLInt64Scalar,
			},
			"nanos": &graphql.InputObjectFieldConfig{
				Type: scalars.GraphQLInt32Scalar,
			},
		}
	}),
})

// Output msg resolver
func ResolveTimestamp(ctx context.Context, i interface{}) (_ *timestamp.Timestamp, rerr error) {
	if i == nil {
		return nil, nil
	}
	args := i.(map[string]interface{})
	_ = args
	var result = new(timestamp.Timestamp)
	// Non-repeated scalar
	if args["seconds"] != nil {
		result.Seconds = args["seconds"].(int64)
	}
	// Non-repeated scalar
	if args["nanos"] != nil {
		result.Nanos = args["nanos"].(int32)
	}

	return result, nil
}

// Maps
// <no value>
func init() {
	// Adding fields to output messages
	// Timestamp message fields
	Timestamp.AddFieldConfig("seconds", &graphql.Field{
		Name: "seconds",
		Type: scalars.GraphQLInt64Scalar,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			src := p.Source.(*timestamp.Timestamp)
			if src == nil {
				return nil, nil
			}
			return src.Seconds, nil
		},
	})
	Timestamp.AddFieldConfig("nanos", &graphql.Field{
		Name: "nanos",
		Type: scalars.GraphQLInt32Scalar,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			src := p.Source.(*timestamp.Timestamp)
			if src == nil {
				return nil, nil
			}
			return src.Nanos, nil
		},
	})
}
