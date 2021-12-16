module gateway

go 1.16

require (
	github.com/alibaba/sentinel-golang v1.0.3
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	go-micro.dev/v4 v4.2.1
	go.opentelemetry.io/contrib v1.1.1
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/trace v1.1.0
	golang.org/x/net v0.0.0-20211109214657-ef0fda0de508 // indirect
	golang.org/x/sys v0.0.0-20211110154304-99a53858aa08 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211020151524-b7c3a969101a
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	pkg v0.0.0
)

replace pkg => ../pkg
