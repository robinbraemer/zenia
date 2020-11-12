COLOR := "\e[1;36m%s\e[0m\n"

##### Compile proto files and generate gapic client for go #####
gen:
	./gen.sh

##### Plugins & tools #####
grpc-install: gogo-protobuf-install
	printf $(COLOR) "Install/update gRPC plugins..."
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

gogo-protobuf-install:
	go get -u github.com/gogo/protobuf/protoc-gen-gogoslick

buf-install:
	printf $(COLOR) "Install/update buf..."
	go get -u github.com/bufbuild/buf/cmd/buf