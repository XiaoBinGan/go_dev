cd Services/protos
protoc --proto_path=. --micro_out=../ --go_out=../ Models.proto
protoc --proto_path=. --micro_out=../ --go_out=../ ProdService.proto
protoc-go-inject-tag -input=../Models.pb.go
cd .. && cd ..