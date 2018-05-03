proto2gql
==============
[![Build Status](https://travis-ci.org/saturn4er/proto2gql.svg?branch=master)](https://travis-ci.org/saturn4er/proto2gql)
[![Coverage Status](https://coveralls.io/repos/github/saturn4er/proto2gql/badge.svg?branch=master)](https://coveralls.io/github/saturn4er/proto2gql?branch=master)

Tool, which generates [graphql-go](https://github.com/graphql-go/graphql) schema for `.proto` file.  

## Installation
```bash
$ go get github.com/saturn4er/proto2gql/cmd/proto2gql
```

## Usage
To generate GraphQL fields by .proto 
```
$ ./proto2gql
```

## Generation process
![Generation process](https://www.dropbox.com/s/9zkutpamcnp0wbb/proto2gql-proc.png)

## Config example
```yaml

paths:                         # path, where parser will search for imports
  - "${GOPATH}/src/"     
generate_tracer: true          # if true, generated code will trace  all functions calls

output_package: "graphql"      # Common Golang package for generated files 
output_path: "./out"           # Path, where generator will put generated files

imports:                       # .proto files imports settings
  output_package: "imports"    # Golang package name for generated imports
  output_path: "./out/imports" # Path, where generator will put generated imports files
  aliases:                     # Global aliases for imports. 
    google/protobuf/timestamp.proto:  "github.com/gogo/protobuf/protobuf/google/protobuf/timestamp.proto"
  settings:
    "${GOPATH}src/github.com/gogo/protobuf/protobuf/google/protobuf/timestamp.proto":
      go_package: "github.com/gogo/protobuf/types"   # golang package, of generated .proto file
      gql_enums_prefix: "TS"                         # prefix, which will be added to all generated GraphQL Enums
      gql_messages_prefix: "TS"                      # prefix, which will be added to all generated GraphQL Messages(including maps)
       

protos:
  - proto_path: "./example/example.proto"           # path to .proto file              
    output_path: "./schema/example"                 # path, where generator will put generated file
    output_package: "example"                       # Golang package for generated file
    paths:                                          # path, where parser will search for imports.  
      - "${GOPATH}/src/github.com/saturn4er/proto2gql/example/"
    gql_messages_prefix: "Example"                  # prefix, which will be added to all generated GraphQL Messages(including maps)
    gql_enums_prefix: "Example"                     # prefix, which will be added to all generated GraphQL Enums
    imports_aliases:                                # imports aliases
      google/protobuf/timestamp.proto:  "github.com/google/protobuf/google/protobuf/timestamp.proto"
    services:             
      ServiceExample:
        alias: "NonServiceExample"                  # service name alias
        methods:  
          queryMethod:                              
            alias: "newQueryMethod"                 # method name alias
            request_type: "QUERY"                   # GraphQL query type (QUERY|MUTATION)
    messages:
      MessageName:
        error_field: "errors"                       # recognize this field as payload error. You can access it in interceptors
        fields:
          message_field: {context_key: "ctx_field_key"}  # Resolver, will try to fetch this field from context instead of fetching it from arguments 
```

## Interceptors

There's two types of Interceptors. The first one can do some logic while parsing GraphQL arguments into request message and the second one, which intercept GRPC call. Here's an example, how to work with it

```go
package main

import (
	"fmt"
	
	"google.golang.org/grpc"
	"github.com/saturn4er/proto2gql/api/interceptors"
)
	

func main(){
    ih := interceptors.InterceptorHandler{}
    ih.OnResolveArgs(func(ctx *interceptors.Context, next interceptors.ResolveArgsInvoker) (result interface{}, err error) {
    	fmt.Println("Before resolving request message")
    	req, err := next()
    	fmt.Println("After resolving request message")
    	return req, err
    })
    ih.OnCall(func(ctx *interceptors.Context, req interface{}, next interceptors.CallMethodInvoker, opts ...grpc.CallOption) (result interface{}, err error) {
        fmt.Println("Before GRPC Call")
        res, err := next(req, opts...)
        fmt.Println("After GRPC Call")
        return res, err
    })
    // queriesFields := GetSomeServiceQueriesFields(someClient, ih)
    // create other schema...
}

```

## How generated code works

![workflow](https://www.dropbox.com/s/at6uuymurq14v2q/proto2gql-execution.png)