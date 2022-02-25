protoc -I/Users/wujiahao/ -I. \
-I ./ \
-I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
--go_out=plugins=grpc:. greeter.proto

protoc -I/Users/wujiahao/ -I. \
-I ./ \
-I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. greeter.proto

