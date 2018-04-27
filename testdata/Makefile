default: proto

proto:
	@protoc -I=${GOPATH}/src:. --go_out=plugins=grpc:.  test.proto
	@protoc -I=${GOPATH}/src:. --go_out=plugins=grpc:.  common/common.proto

.PHONY: proto
