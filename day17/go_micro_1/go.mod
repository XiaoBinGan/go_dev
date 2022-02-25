module go_dev/day17/go_micro_1

go 1.13

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/micro/protoc-gen-micro v1.0.0 // indirect
	golang.org/x/net v0.0.0-20210222171744-9060382bd457 // indirect
	golang.org/x/sys v0.0.0-20210220050731-9a76102bfb43 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210222212404-3e1e516060db // indirect
	google.golang.org/grpc v1.35.0 // indirect
	google.golang.org/grpc/examples v0.0.0-20210223000434-25cf9393fa21 // indirect
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
