default: install

install:
	@go install ./cmd/proto2gql/

.PHONY:
	install
