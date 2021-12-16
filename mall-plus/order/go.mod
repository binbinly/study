module order

go 1.16

require (
	common v0.0.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/pkg/errors v0.9.1
	github.com/streadway/amqp v1.0.0
	go-micro.dev/v4 v4.2.1
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.21.15
	pkg v0.0.0
)

replace pkg => ../pkg

replace common => ./../common
