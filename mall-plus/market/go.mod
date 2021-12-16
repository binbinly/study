module market

go 1.16

require (
	common v0.0.0
	github.com/go-redis/redis/v8 v8.11.3
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	go-micro.dev/v4 v4.2.1
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.21.15
	pkg v0.0.0
)

replace pkg => ../pkg

replace common => ./../common
