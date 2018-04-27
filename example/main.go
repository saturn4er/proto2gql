package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/graphql-go/graphql"
	"github.com/saturn4er/proto2gql/example/out/example"
	proto "github.com/saturn4er/proto2gql/example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FakeClient struct {
}

func (FakeClient) GetEmptiesMsg(ctx context.Context, in *timestamp.Timestamp, opts ...grpc.CallOption) (*proto.Empty, error) {
	panic("implement me")
}

func (FakeClient) GetQueryMethod(ctx context.Context, in *proto.AOneOffs, opts ...grpc.CallOption) (*proto.B, error) {
	return nil, nil
}

func (FakeClient) MutationMethod(ctx context.Context, in *proto.B, opts ...grpc.CallOption) (*proto.A, error) {
	return nil, nil
}

func (FakeClient) QueryMethod(ctx context.Context, in *proto.A, opts ...grpc.CallOption) (*proto.B, error) {
	return nil, nil
}

func (FakeClient) GetMutatuionMethod(ctx context.Context, in *proto.MsgWithEmpty, opts ...grpc.CallOption) (*proto.MsgWithEmpty, error) {
	return nil, nil
}

func main() {
	// Schema
	fieldsQuery := example.GetServiceExampleGraphQLQueriesFields(new(FakeClient), nil)
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fieldsQuery}
	fieldsMutation := example.GetServiceExampleGraphQLMutationsFields(new(FakeClient), nil)
	rootMutation := graphql.ObjectConfig{Name: "Mutation", Fields: fieldsMutation}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(rootMutation)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			getQueryMethod{
				n_r_enum
			}
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
