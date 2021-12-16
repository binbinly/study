module third-party

go 1.16

require (
	common v0.0.0
	github.com/asim/go-micro/v3 v3.6.0
	github.com/ethereum/go-ethereum v1.10.9
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.27.1
	pkg v0.0.0
)

replace pkg => ../pkg

replace common => ./../common
