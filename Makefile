
gen:
	protoc --experimental_allow_proto3_optional --proto_path=${GOPATH}/src:. --twirp_out=. --go_out=. ./rpc/interceptor/interceptor.proto