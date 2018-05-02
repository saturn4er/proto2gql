default: install

install:
	@go install ./cmd/proto2gql/

build:
	@go build -o ./bin/proto2gql ./cmd/proto2gql

test:
	go test ./...


.PHONY: install testdata
