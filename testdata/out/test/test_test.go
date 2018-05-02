package test

import (
	"github.com/graphql-go/graphql"
	"github.com/saturn4er/proto2gql/api/interceptors"
	"github.com/saturn4er/proto2gql/testdata"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

type TestMock struct {
	mock.Mock
}

var _ testdata.ServiceExampleClient = TestMock{}

func (TestMock) GetQueryMethod(ctx context.Context, in *testdata.RootMessage, opts ...grpc.CallOption) (*testdata.RootMessage2, error) {
	panic("implement me")
}

func (TestMock) MutationMethod(ctx context.Context, in *testdata.RootMessage2, opts ...grpc.CallOption) (*testdata.RootMessage_NestedMessage, error) {
	panic("implement me")
}

func (TestMock) EmptyMsgs(ctx context.Context, in *testdata.Empty, opts ...grpc.CallOption) (*testdata.Empty, error) {
	panic("implement me")
}

func (TestMock) MsgsWithEpmty(ctx context.Context, in *testdata.MessageWithEmpty, opts ...grpc.CallOption) (*testdata.MessageWithEmpty, error) {
	panic("implement me")
}

type interceptions struct {
	mock.Mock
}

func (i *interceptions) resolveArgs(ctx *interceptors.Context, next interceptors.ResolveArgsInvoker) (result interface{}, err error) {
	args := i.Called(ctx, next)
	return args.Get(0), args.Error(1)
}

func (i *interceptions) call(ctx *interceptors.Context, req interface{}, next interceptors.CallMethodInvoker, opts ...grpc.CallOption) (result interface{}, err error) {
	args := i.Called(ctx, req, next, opts)
	return args.Get(0), args.Error(1)
}

func TestFields(t *testing.T) {
	Convey("Test example fields", t, func() {
		client := TestMock{}
		Convey("Test interceptors", func() {
			ih := new(interceptors.InterceptorHandler)
			is := new(interceptions)
			ih.OnResolveArgs(is.resolveArgs)
			ih.OnCall(is.call)
			schema, err := GetSchema(client, ih)
			So(err, ShouldBeNil)
			graphql.Do(graphql.Params{
				Schema:        schema,
				RequestString: "",
			})
		})
	})
}

func GetSchema(client testdata.ServiceExampleClient, handler *interceptors.InterceptorHandler) (graphql.Schema, error) {
	queries := GetServiceExampleGraphQLQueriesFields(client, handler)
	mutations := GetServiceExampleGraphQLMutationsFields(client, handler)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queries}
	rootMutation := graphql.ObjectConfig{Name: "Mutation", Fields: mutations}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(rootMutation)}
	return graphql.NewSchema(schemaConfig)
}
