proto2gql
==============
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

## Config example
```yaml

paths:
  - "${GOPATH}/src/" # path, where tool will search for imports
generate_tracer: true 

output_package: "graphql"
output_path: "./out"

imports:
  output_package: "imports"
  output_path: "./out/imports" # Where to put generated imports
  aliases:
    google/protobuf/timestamp.proto:  "github.com/gogo/protobuf/protobuf/google/protobuf/timestamp.proto"
  settings:
    "${GOPATH}src/github.com/gogo/protobuf/protobuf/google/protobuf/timestamp.proto":
      go_package: "github.com/gogo/protobuf/types"

protos:
  - proto_path: "./example/example.proto"              
    output_path: "./schema/example"         
    output_package: "example"
    paths:
      - "${GOPATH}/src/github.com/saturn4er/proto2gql/example/"
    gql_messages_prefix: "Example"
    gql_enums_prefix: "Example"
    services:
      ServiceExample:
        methods:
          queryMethod:
            alias: "newQueryMethod"
            request_type: "QUERY"
    messages:
      A:
        fields:
          map_enum: {context_key: "map_enum_from_context"}  # will try to fetch this field from context instead of fetching it from user 
```