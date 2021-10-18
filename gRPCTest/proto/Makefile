go:
	protoc -I . -I $$GOPATH/src --go_out=. *.proto
cpp:
	protoc -I . -I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --cpp_out=:. *.proto

