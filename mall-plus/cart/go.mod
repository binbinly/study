module cart

go 1.16

require (
	common v0.0.0
	github.com/go-redis/redis/v8 v8.11.3
	github.com/pkg/errors v0.9.1
	go-micro.dev/v4 v4.2.1
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	google.golang.org/protobuf v1.27.1
	pkg v0.0.0
)

replace pkg => ../pkg

replace common => ./../common
