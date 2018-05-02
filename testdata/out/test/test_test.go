package test

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/location"
	"github.com/pkg/errors"
	"github.com/saturn4er/proto2gql/api/interceptors"
	"github.com/saturn4er/proto2gql/testdata"
	"github.com/saturn4er/proto2gql/testdata/common"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

type testCase struct {
	Name           string
	ResolveArgsRes interface{}
	ResolveArgsErr error
	CallRes        interface{}
	CallErr        error
	ExpectedResult interface{}
	ErrorResult    bool
}
type TestMock struct {
	mock.Mock
}

var _ testdata.ServiceExampleClient = new(TestMock)

func (t *TestMock) GetQueryMethod(ctx context.Context, in *testdata.RootMessage, opts ...grpc.CallOption) (*testdata.RootMessage, error) {
	res := t.Called(ctx, in, opts)
	return res.Get(0).(*testdata.RootMessage), res.Error(1)
}

func (t *TestMock) MutationMethod(ctx context.Context, in *testdata.RootMessage2, opts ...grpc.CallOption) (*testdata.RootMessage_NestedMessage, error) {
	res := t.Called(ctx, in, opts)
	return res.Get(0).(*testdata.RootMessage_NestedMessage), res.Error(1)
}

func (t *TestMock) EmptyMsgs(ctx context.Context, in *testdata.Empty, opts ...grpc.CallOption) (*testdata.Empty, error) {
	res := t.Called(ctx, in, opts)
	return res.Get(0).(*testdata.Empty), res.Error(1)
}

func (t *TestMock) MsgsWithEpmty(ctx context.Context, in *testdata.MessageWithEmpty, opts ...grpc.CallOption) (*testdata.MessageWithEmpty, error) {
	res := t.Called(ctx, in, opts)
	return res.Get(0).(*testdata.MessageWithEmpty), res.Error(1)
}

type interceptions struct {
	mock.Mock
}

func (i *interceptions) resolveArgs(ctx *interceptors.Context, next interceptors.ResolveArgsInvoker) (result interface{}, err error) {
	args := i.Called(ctx, next)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	if res != nil {
		return res, nil
	}
	return next()
}

func (i *interceptions) call(ctx *interceptors.Context, req interface{}, next interceptors.CallMethodInvoker, opts ...grpc.CallOption) (result interface{}, err error) {
	args := i.Called(ctx, req, next, opts)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	if res != nil {
		return res, nil
	}
	return next(req, opts...)
}
func TestMutationMethod(t *testing.T) {
	Convey("Setup", t, func() {
		client := new(TestMock)
		ih := new(interceptors.InterceptorHandler)
		validResult := &graphql.Result{
			Errors: nil,
			Data: map[string]interface{}{
				"mutationMethod": map[string]interface{}{
					"sub_r_enum":     []interface{}{"NestedEnumVal1"},
					"sub_sub_r_enum": []interface{}{"NestedNestedEnumVal1", "NestedNestedEnumVal2"},
				},
			},
		}
		client.On("MutationMethod", mock.Anything, mock.Anything, mock.MatchedBy(func(val []grpc.CallOption) bool {
			return len(val) == 0
		})).Return(&testdata.RootMessage_NestedMessage{
			SubREnum: []testdata.RootMessage_NestedEnum{testdata.RootMessage_NestedEnumVal1},
			SubSubREnum: []testdata.RootMessage_NestedMessage_NestedNestedEnum{
				testdata.RootMessage_NestedMessage_NestedNestedEnumVal1,
				testdata.RootMessage_NestedMessage_NestedNestedEnumVal2,
			},
		}, nil)

		Convey("With interceptors", func() {
			is := new(interceptions)
			ih.OnResolveArgs(is.resolveArgs)
			ih.OnCall(is.call)
			schema, err := GetSchema(client, ih)
			So(err, ShouldBeNil)

			var ErrResult = func(errMsg string) *graphql.Result {
				return &graphql.Result{
					Data: nil,
					Errors: []gqlerrors.FormattedError{
						{Message: errMsg, Locations: []location.SourceLocation{}},
					},
				}
			}
			var tests = []testCase{
				{"Positive flow", nil, nil, nil, nil, validResult, false},
				{"Resolve args err", nil, errors.New("some err"), nil, nil, ErrResult("some err"), true},
				{"Resolve args bad value", new(int), nil, nil, nil, ErrResult("Resolve args interceptor returns bad request type(*int). Should be: *testdata.RootMessage2"), true},
				{"Call bad value", nil, nil, new(int), nil, ErrResult("Call Interceptor returns bad value type(*int). Should return *testdata.RootMessage_NestedMessage"), true},
			}
			for _, t := range tests {
				Convey(t.Name, func() {
					is.On("resolveArgs", mock.Anything, mock.Anything).Return(t.ResolveArgsRes, t.ResolveArgsErr).Once()
					is.On("call", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(t.CallRes, t.CallErr).Once()
					res := graphql.Do(graphql.Params{
						Schema:        schema,
						RequestString: "mutation { mutationMethod {sub_r_enum, sub_sub_r_enum} }",
					})
					So(res, ShouldResemble, t.ExpectedResult)
				})

			}
		})
		Convey("Without interceptors", func() {
			schema, err := GetSchema(client, nil)
			So(err, ShouldBeNil)
			res := graphql.Do(graphql.Params{
				Schema:        schema,
				RequestString: "mutation { mutationMethod {sub_r_enum, sub_sub_r_enum} }",
			})
			So(res, ShouldResemble, validResult)
		})
	})
}
func TestMsgWithEmpty(t *testing.T) {
	Convey("Setup", t, func() {
		client := new(TestMock)
		ih := new(interceptors.InterceptorHandler)
		validResult := &graphql.Result{
			Errors: nil,
			Data: map[string]interface{}{
				"MsgsWithEpmty": map[string]interface{}{
					"empt": nil,
				},
			},
		}
		client.On("MsgsWithEpmty", mock.Anything, mock.Anything, mock.MatchedBy(func(val []grpc.CallOption) bool {
			return len(val) == 0
		})).Return(&testdata.MessageWithEmpty{Empt: nil}, nil)

		Convey("With interceptors", func() {
			is := new(interceptions)
			ih.OnResolveArgs(is.resolveArgs)
			ih.OnCall(is.call)
			schema, err := GetSchema(client, ih)
			So(err, ShouldBeNil)

			var ErrResult = func(errMsg string) *graphql.Result {
				return &graphql.Result{
					Data: nil,
					Errors: []gqlerrors.FormattedError{
						{Message: errMsg, Locations: []location.SourceLocation{}},
					},
				}
			}
			var tests = []testCase{
				{"Positive flow", nil, nil, nil, nil, validResult, false},
				{"Resolve args err", nil, errors.New("some err"), nil, nil, ErrResult("some err"), true},
				{"Resolve args bad value", new(int), nil, nil, nil, ErrResult("Resolve args interceptor returns bad request type(*int). Should be: *testdata.MessageWithEmpty"), true},
				{"Call bad value", nil, nil, new(int), nil, ErrResult("Call Interceptor returns bad value type(*int). Should return *testdata.MessageWithEmpty"), true},
			}
			for _, t := range tests {
				Convey(t.Name, func() {
					is.On("resolveArgs", mock.Anything, mock.Anything).Return(t.ResolveArgsRes, t.ResolveArgsErr).Once()
					is.On("call", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(t.CallRes, t.CallErr).Once()
					res := graphql.Do(graphql.Params{
						Schema:        schema,
						RequestString: "mutation { MsgsWithEpmty {empt} }",
					})
					So(res, ShouldResemble, t.ExpectedResult)
				})

			}
		})
		Convey("Without interceptors", func() {
			schema, err := GetSchema(client, nil)
			So(err, ShouldBeNil)
			res := graphql.Do(graphql.Params{
				Schema:        schema,
				RequestString: "mutation { MsgsWithEpmty {empt} }",
			})
			So(res, ShouldResemble, validResult)
		})
	})
}
func TestEmptyMsgs(t *testing.T) {
	Convey("Setup", t, func() {
		client := new(TestMock)
		ih := new(interceptors.InterceptorHandler)
		var validResult = &graphql.Result{
			Data:   map[string]interface{}{"EmptyMsgs": nil},
			Errors: nil,
		}

		client.On("EmptyMsgs", mock.Anything, mock.Anything, mock.MatchedBy(func(val []grpc.CallOption) bool {
			return len(val) == 0
		})).Return(new(testdata.Empty), nil)

		Convey("With interceptors", func() {
			is := new(interceptions)
			ih.OnResolveArgs(is.resolveArgs)
			ih.OnCall(is.call)
			schema, err := GetSchema(client, ih)
			So(err, ShouldBeNil)
			var ErrResult = func(errMsg string) *graphql.Result {
				return &graphql.Result{
					Data: nil,
					Errors: []gqlerrors.FormattedError{
						{Message: errMsg, Locations: []location.SourceLocation{}},
					},
				}
			}
			var tests = []testCase{
				{"Positive flow", nil, nil, nil, nil, validResult, false},
				{"Resolve args err", nil, errors.New("some err"), nil, nil, ErrResult("some err"), true},
				{"Resolve args bad value", new(int), nil, nil, nil, ErrResult("Resolve args interceptor returns bad request type(*int). Should be: *testdata.Empty"), true},
				{"Call bad value", nil, nil, new(int), nil, ErrResult("Call Interceptor returns bad value type(*int). Should return *testdata.Empty"), true},
			}
			for _, t := range tests {
				Convey(t.Name, func() {
					is.On("resolveArgs", mock.Anything, mock.Anything).Return(t.ResolveArgsRes, t.ResolveArgsErr).Once()
					is.On("call", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(t.CallRes, t.CallErr).Once()
					res := graphql.Do(graphql.Params{
						Schema:        schema,
						RequestString: "mutation { EmptyMsgs }",
					})
					So(res, ShouldResemble, t.ExpectedResult)
				})

			}
		})
		Convey("Without interceptors", func() {
			schema, err := GetSchema(client, nil)
			So(err, ShouldBeNil)
			res := graphql.Do(graphql.Params{
				Schema:        schema,
				RequestString: "mutation { EmptyMsgs }",
			})
			So(res, ShouldResemble, validResult)
		})
	})
}
func TestGetQueryMethod(t *testing.T) {
	Convey("Setup", t, func() {
		client := new(TestMock)
		ih := new(interceptors.InterceptorHandler)
		var nestedMessage = &testdata.RootMessage_NestedMessage{
			SubREnum: []testdata.RootMessage_NestedEnum{
				testdata.RootMessage_NestedEnumVal1,
				testdata.RootMessage_NestedEnumVal0,
			},
			SubSubREnum: []testdata.RootMessage_NestedMessage_NestedNestedEnum{
				testdata.RootMessage_NestedMessage_NestedNestedEnumVal2,
				testdata.RootMessage_NestedMessage_NestedNestedEnumVal0,
				testdata.RootMessage_NestedMessage_NestedNestedEnumVal1,
			},
		}

		var validResult = &graphql.Result{
			Data: map[string]interface{}{
				"getQueryMethod": map[string]interface{}{
					"map_enum": []interface{}{
						map[string]interface{}{"key": int32(1), "value": nil},
						map[string]interface{}{"key": int32(10), "value": nil},
					},
					"map_scalar": []interface{}{
						map[string]interface{}{"key": int32(10), "value": int32(20)},
						map[string]interface{}{"key": int32(30), "value": int32(40)},
					},
					"map_msg": []interface{}{
						map[string]interface{}{
							"key": int32(30),
							"value": map[string]interface{}{
								"sub_r_enum":     []interface{}{"NestedEnumVal1", "NestedEnumVal0"},
								"sub_sub_r_enum": []interface{}{"NestedNestedEnumVal2", "NestedNestedEnumVal0", "NestedNestedEnumVal1"},
							},
						},
					},
					"e_f_o_e":       "CommonEnumVal0",
					"e_f_o_em":      nil,
					"e_f_o_m":       map[string]interface{}{"scalar": int32(123)},
					"e_f_o_s":       int32(0),
					"em_f_o_em":     nil,
					"em_f_o_en":     "RootEnumVal0",
					"em_f_o_m":      map[string]interface{}{"some_field": nil},
					"em_f_o_s":      int32(0),
					"m_f_o_e":       "RootEnumVal2",
					"m_f_o_em":      nil,
					"m_f_o_m":       map[string]interface{}{"some_field": nil},
					"m_f_o_s":       int32(0),
					"n_r_empty_msg": nil,
					"n_r_enum":      "CommonEnumVal0",
					"n_r_msg":       map[string]interface{}{"scalar": int32(1532)},
					"n_r_scalar":    int32(56123),
					"r_empty_msg":   []interface{}{nil, nil, nil},
					"r_enum":        []interface{}{"RootEnumVal1", "RootEnumVal2"},
					"r_msg": []interface{}{
						map[string]interface{}{
							"sub_r_enum":     []interface{}{"NestedEnumVal1", "NestedEnumVal0"},
							"sub_sub_r_enum": []interface{}{"NestedNestedEnumVal2", "NestedNestedEnumVal0", "NestedNestedEnumVal1"},
						},
						map[string]interface{}{
							"sub_r_enum":     []interface{}{"NestedEnumVal1", "NestedEnumVal0"},
							"sub_sub_r_enum": []interface{}{"NestedNestedEnumVal2", "NestedNestedEnumVal0", "NestedNestedEnumVal1"},
						},
					},
					"r_scalar":            []interface{}{int32(1), int32(2), int32(3), int32(4)},
					"s_f_o_e":             "RootEnumVal0",
					"s_f_o_m":             nil,
					"s_f_o_mes":           map[string]interface{}{"some_field": int32(123)},
					"s_f_o_s":             int32(0),
					"scalar_from_context": int32(144123253),
				},
			},
			Errors: nil,
		}
		checkValidResponse := func(res *graphql.Result) {
			So(res.Errors, ShouldBeNil)
			So(res.Data, ShouldNotBeNil)
			var resultFields = res.Data.(map[string]interface{})["getQueryMethod"].(map[string]interface{})
			var validResultFields = validResult.Data.(map[string]interface{})["getQueryMethod"].(map[string]interface{})
			for _, value := range []string{"r_msg", "r_scalar", "r_enum", "r_empty_msg", "n_r_enum", "n_r_scalar", "n_r_msg", "scalar_from_context", "n_r_empty_msg", "e_f_o_e", "e_f_o_s", "e_f_o_m", "e_f_o_em", "s_f_o_s", "s_f_o_e", "s_f_o_mes", "s_f_o_m", "m_f_o_m", "m_f_o_s", "m_f_o_e", "m_f_o_em", "em_f_o_em", "em_f_o_s", "em_f_o_en", "em_f_o_m"} {
				Convey(value+" field should be valid", func() {
					So(resultFields[value], ShouldResemble, validResultFields[value])
				})
			}
			for _, value := range []string{"map_enum", "map_scalar", "map_msg"} {
				Convey(value+" field should be valid", func() {
					resList := resultFields[value].([]interface{})
					validResList := validResultFields[value].([]interface{})
					So(resList, ShouldHaveLength, len(validResList))

					for _, vf := range validResList {
						var found bool
						for _, rf := range resList {
							validResVal := vf.(map[string]interface{})
							resVal := rf.(map[string]interface{})
							if resVal["key"] == validResVal["key"] {
								So(resVal, ShouldResemble, validResVal)
								found = true
							}
						}
						So(found, ShouldBeTrue)
					}
				})
			}
		}

		Convey("With interceptors", func() {
			client.On("GetQueryMethod", mock.Anything, mock.Anything, mock.MatchedBy(func(val []grpc.CallOption) bool {
				return len(val) == 0
			})).Return(&testdata.RootMessage{
				MapEnum: map[int32]testdata.RootMessage_NestedEnum{
					1:  testdata.RootMessage_NestedEnumVal1,
					10: testdata.RootMessage_NestedEnumVal0,
				},
				MapScalar: map[int32]int32{
					10: 20,
					30: 40,
				},
				MapMsg: map[int32]*testdata.RootMessage_NestedMessage{
					30: nestedMessage,
				},
				RMsg:               []*testdata.RootMessage_NestedMessage{nestedMessage, nestedMessage},
				RScalar:            []int32{1, 2, 3, 4},
				REnum:              []testdata.RootEnum{testdata.RootEnum_RootEnumVal1, testdata.RootEnum_RootEnumVal2},
				REmptyMsg:          []*testdata.Empty{{}, {}, {}},
				NREnum:             common.CommonEnum_CommonEnumVal0,
				NRScalar:           56123,
				NRMsg:              &common.CommonMessage{Scalar: 1532},
				ScalarFromContext:  144123253,
				NREmptyMsg:         new(testdata.Empty),
				EnumFirstOneoff:    &testdata.RootMessage_EFOM{&common.CommonMessage{123}},
				ScalarFirstOneoff:  &testdata.RootMessage_SFOMes{&testdata.RootMessage2{123}},
				MessageFirstOneoff: &testdata.RootMessage_MFOE{testdata.RootEnum_RootEnumVal2},
				EmptyFirstOneoff:   &testdata.RootMessage_EmFOEn{testdata.RootEnum_RootEnumVal0},
			}, nil)
			is := new(interceptions)
			ih.OnResolveArgs(is.resolveArgs)
			ih.OnCall(is.call)
			schema, err := GetSchema(client, ih)
			So(err, ShouldBeNil)
			var ErrResult = func(errMsg string) *graphql.Result {
				return &graphql.Result{
					Data: nil,
					Errors: []gqlerrors.FormattedError{
						{Message: errMsg, Locations: []location.SourceLocation{}},
					},
				}
			}
			var tests = []testCase{
				{"Positive flow", nil, nil, nil, nil, nil, false},
				{"Resolve args err", nil, errors.New("some err"), nil, nil, ErrResult("some err"), true},
				{"Resolve args bad value", new(int), nil, nil, nil, ErrResult("Resolve args interceptor returns bad request type(*int). Should be: *testdata.RootMessage"), true},
				{"Call bad value", nil, nil, new(int), nil, ErrResult("Call Interceptor returns bad value type(*int). Should return *testdata.RootMessage"), true},
			}
			for _, t := range tests {
				Convey(t.Name, func() {
					is.On("resolveArgs", mock.Anything, mock.Anything).Return(t.ResolveArgsRes, t.ResolveArgsErr).Once()
					is.On("call", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(t.CallRes, t.CallErr).Once()
					res := graphql.Do(graphql.Params{
						Schema: schema,
						RequestString: `query { 
									getQueryMethod {
										map_enum{key,value}
										map_scalar{key,value}
										map_msg{
											key
											value{
												sub_r_enum 
												sub_sub_r_enum
											}
										}
										r_msg{sub_r_enum, sub_sub_r_enum}
										r_scalar
										r_enum
										r_empty_msg
										n_r_enum
										n_r_scalar
										n_r_msg{scalar}
										scalar_from_context
										n_r_empty_msg
										e_f_o_e
										e_f_o_s
										e_f_o_m{scalar}
										e_f_o_em
										s_f_o_s
										s_f_o_e
										s_f_o_mes{some_field}
										s_f_o_m
										m_f_o_m{some_field}
										m_f_o_s
										m_f_o_e
										m_f_o_em
										em_f_o_em
										em_f_o_s
										em_f_o_en
										em_f_o_m{some_field}
									}
								}`,
						Context: context.Background(),
					})
					if t.ErrorResult {
						So(res, ShouldResemble, t.ExpectedResult)
					} else {
						checkValidResponse(res)
					}
				})

			}
		})
		Convey("Without interceptors", func() {
			schema, err := GetSchema(client, nil)
			client.On("GetQueryMethod", mock.Anything, mock.Anything, mock.MatchedBy(func(val []grpc.CallOption) bool {
				return len(val) == 0
			})).Return(&testdata.RootMessage{
				MapEnum: map[int32]testdata.RootMessage_NestedEnum{
					1:  testdata.RootMessage_NestedEnumVal1,
					10: testdata.RootMessage_NestedEnumVal0,
				},
				MapScalar: map[int32]int32{
					10: 20,
					30: 40,
				},
				MapMsg: map[int32]*testdata.RootMessage_NestedMessage{
					30: nestedMessage,
				},
				RMsg:               []*testdata.RootMessage_NestedMessage{nestedMessage, nestedMessage},
				RScalar:            []int32{1, 2, 3, 4},
				REnum:              []testdata.RootEnum{testdata.RootEnum_RootEnumVal1, testdata.RootEnum_RootEnumVal2},
				REmptyMsg:          []*testdata.Empty{{}, {}, {}},
				NREnum:             common.CommonEnum_CommonEnumVal0,
				NRScalar:           56123,
				NRMsg:              &common.CommonMessage{Scalar: 1532},
				ScalarFromContext:  144123253,
				NREmptyMsg:         new(testdata.Empty),
				EnumFirstOneoff:    &testdata.RootMessage_EFOM{&common.CommonMessage{123}},
				ScalarFirstOneoff:  &testdata.RootMessage_SFOMes{&testdata.RootMessage2{123}},
				MessageFirstOneoff: &testdata.RootMessage_MFOE{testdata.RootEnum_RootEnumVal2},
				EmptyFirstOneoff:   &testdata.RootMessage_EmFOEn{testdata.RootEnum_RootEnumVal0},
			}, nil)
			So(err, ShouldBeNil)

			res := graphql.Do(graphql.Params{
				Schema: schema,
				RequestString: `query { 
									getQueryMethod {
										map_enum{key,value}
										map_scalar{key,value}
										map_msg{
											key
											value{
												sub_r_enum 
												sub_sub_r_enum
											}
										}
										r_msg{sub_r_enum, sub_sub_r_enum}
										r_scalar
										r_enum
										r_empty_msg
										n_r_enum
										n_r_scalar
										n_r_msg{scalar}
										scalar_from_context
										n_r_empty_msg
										e_f_o_e
										e_f_o_s
										e_f_o_m{scalar}
										e_f_o_em
										s_f_o_s
										s_f_o_e
										s_f_o_mes{some_field}
										s_f_o_m
										m_f_o_m{some_field}
										m_f_o_s
										m_f_o_e
										m_f_o_em
										em_f_o_em
										em_f_o_s
										em_f_o_en
										em_f_o_m{some_field}
									}
								}`,
				Context: context.WithValue(context.Background(), "ctx_key", int32(123)),
			})
			checkValidResponse(res)
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
