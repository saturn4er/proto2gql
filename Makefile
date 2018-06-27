default: build_templates install

install:
	@go install ./cmd/proto2gql/

build:
	@go build -o ./bin/proto2gql ./cmd/proto2gql

build_templates:
	go-bindata -prefix ./generator/common -o ./generator/common/templates.go -pkg common ./generator/common/templates
	go-bindata -prefix ./generator/schema -o ./generator/schema/templates.go -pkg schema ./generator/schema/templates

test:
	go test ./...


.PHONY: install


