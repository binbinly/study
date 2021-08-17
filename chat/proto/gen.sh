protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative base/base.proto
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative connect/connect.proto
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative chat/chat.proto
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative center/center.proto