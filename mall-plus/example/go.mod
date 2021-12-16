module example

go 1.16

require (
	github.com/asim/go-micro/plugins/broker/rabbitmq/v4 v4.0.0-20211103025805-c5be9f560cdb
	github.com/asim/go-micro/plugins/client/http/v4 v4.0.0-20211015130340-9f4770e7fdfc
	github.com/asim/go-micro/plugins/config/encoder/yaml/v4 v4.0.0-20211111140334-799b8d6a6559
	github.com/asim/go-micro/plugins/registry/memory/v4 v4.0.0-20211014070921-e7dbda689e97 // indirect
	github.com/asim/go-micro/plugins/server/http/v4 v4.0.0-20211015130340-9f4770e7fdfc
	github.com/gin-gonic/gin v1.7.4
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/streadway/amqp v1.0.0
	go-micro.dev/v4 v4.2.1
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	google.golang.org/genproto v0.0.0-20211013025323-ce878158c4d4
	google.golang.org/grpc v1.40.0
	google.golang.org/grpc/examples v0.0.0-20210902184326-c93e472777b9
	google.golang.org/protobuf v1.27.1
	pkg v0.0.0
)

replace pkg => ../pkg
