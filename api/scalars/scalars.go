package scalars

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

var GraphQLInt64Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Int64",
	Description: "The `Int64` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1. \n" +
		"Can be passed like a string",
	Serialize: func(value interface{}) interface{} {
		switch val := value.(type) {
		case int64:
			return val
		case int32:
			return int64(val)
		case int:
			return int64(val)
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch val := value.(type) {
		case string:
			value, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil
			}
			return value
		case int32:
			return int64(val)
		case int64:
			return val
		case float32:
			return int64(val)
		case float64:
			return int64(val)
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.GetKind() {
		case kinds.IntValue, kinds.StringValue:
			val, err := strconv.ParseInt(valueAST.GetValue().(string), 10, 64)
			if err != nil {
				return nil
			}
			return val
		}
		return nil
	},
})

var GraphQLInt32Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Int32",
	Description: "The  `Int32` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1. \n" +
		"Can be passed like a string",
	Serialize: func(value interface{}) interface{} {
		switch val := value.(type) {
		case int32:
			return val
		case int:
			return int32(val)
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch val := value.(type) {
		case string:
			value, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return nil
			}
			return int32(value)
		case int32:
			return value
		case int64:
			return int32(val)
		case float32:
			return int32(val)
		case float64:
			return int32(val)
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.GetKind() {
		case kinds.IntValue, kinds.StringValue:
			val, err := strconv.ParseInt(valueAST.GetValue().(string), 10, 32)
			if err != nil {
				return nil
			}
			return int32(val)
		}
		return nil
	},
})

var GraphQLUInt64Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "UInt64",
	Description: "The `UInt64` scalar type represents non-fractional unsigned whole numeric values. Int can represent values between 0 and 2^64 - 1.\n" +
		"Can be passed like a string",
	Serialize: func(value interface{}) interface{} {
		switch val := value.(type) {
		case uint64:
			return val
		case uint32:
			return uint64(val)
		case uint:
			return uint64(val)
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch val := value.(type) {
		case string:
			value, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return nil
			}
			return value
		case uint32:
			return uint64(val)
		case uint64:
			return val
		case float32:
			return uint64(val)
		case float64:
			return uint64(val)
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.GetKind() {
		case kinds.IntValue, kinds.StringValue:
			val, err := strconv.ParseUint(valueAST.GetValue().(string), 10, 64)
			if err != nil {
				return nil
			}
			return val
		}
		return nil
	},
})

var GraphQLUInt32Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "UInt32",
	Description: "The `UInt32` scalar type represents non-fractional unsigned whole numeric values. Int can represent values between 0 and 2^32 - 1.\n" +
		"Can be passed like a string",
	Serialize: func(value interface{}) interface{} {
		switch val := value.(type) {
		case uint32:
			return val
		case uint:
			return uint32(val)
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch val := value.(type) {
		case string:
			value, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				return nil
			}
			return uint32(value)
		case uint32:
			return val
		case float32:
			return uint32(val)
		case float64:
			return uint32(val)
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.GetKind() {
		case kinds.IntValue, kinds.StringValue:
			val, err := strconv.ParseUint(valueAST.GetValue().(string), 10, 32)
			if err != nil {
				return nil
			}
			return uint32(val)
		}
		return nil
	},
})

var GraphQLFloat32Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Float32",
	Serialize: func(value interface{}) interface{} {
		if val, ok := value.(float32); ok {
			return val
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch val := value.(type) {
		case string:
			value, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return nil
			}
			return float32(value)
		case float64:
			return float32(val)
		case float32:
			return val
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.GetKind() {

		case kinds.IntValue, kinds.StringValue, kinds.FloatValue:
			val, err := strconv.ParseFloat(valueAST.GetValue().(string), 32)
			if err != nil {
				return nil
			}
			return float32(val)
		}
		return nil
	},
})

var GraphQLFloat64Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Float64",
	Serialize: func(value interface{}) interface{} {
		switch val := value.(type) {
		case float32:
			return float64(val)
		case float64:
			return val
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch val := value.(type) {
		case string:
			value, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil
			}
			return float64(value)
		case float32:
			return float64(val)
		case float64:
			return val
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.GetKind() {
		case kinds.IntValue, kinds.StringValue, kinds.FloatValue:
			val, err := strconv.ParseFloat(valueAST.GetValue().(string), 64)
			if err != nil {
				return nil
			}
			return float64(val)
		}

		return nil
	},
})
var NoDataScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "NoData",
	Description: "The `NoData` scalar type represents no data.",
	Serialize: func(value interface{}) interface{} {
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		return 0
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return 0
	},
})
