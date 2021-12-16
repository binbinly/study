module common

go 1.16

require (
	github.com/asim/go-micro/plugins/client/grpc/v4 v4.0.0-20211019191242-9edc569e68bb
	github.com/asim/go-micro/plugins/config/encoder/yaml/v4 v4.0.0-20211014070921-e7dbda689e97
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20211014070921-e7dbda689e97
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211014070921-e7dbda689e97
	github.com/asim/go-micro/plugins/server/grpc/v4 v4.0.0-20211023082042-af3cfa0a4cac
	github.com/asim/go-micro/plugins/wrapper/breaker/gobreaker/v4 v4.0.0-20211111140334-799b8d6a6559
	github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v4 v4.0.0-20211111140334-799b8d6a6559
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4 v4.0.0-20211111140334-799b8d6a6559
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.3.1
	github.com/urfave/cli/v2 v2.3.0
	go-micro.dev/v4 v4.2.1
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/trace v1.1.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.21.15
	pkg v0.0.0
)

replace pkg => ../pkg
