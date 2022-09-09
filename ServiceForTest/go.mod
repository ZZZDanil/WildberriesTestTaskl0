module ggg

go 1.14

replace service => ../Service

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/nats-io/nats-server/v2 v2.8.4 // indirect
	github.com/nats-io/nats-streaming-server v0.24.6 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	service v0.0.0-00010101000000-000000000000
)
