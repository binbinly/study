module product

go 1.16

require (
	common v0.0.0
	github.com/golang/glog v1.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/olivere/elastic/v7 v7.0.29
	github.com/pkg/errors v0.9.1
	go-micro.dev/v4 v4.2.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211004093028-2c5d950f24ef // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.21.15
	pkg v0.0.0
)

replace pkg => ../pkg

replace common => ./../common
