package main

import (
	"chat-micro/cmd"
)

// @title chat docs api
// @version 1.0
// @description chat api

// @contact.name test
// @contact.url http://www.swagger.io/support
// @contact.email test@test.com

// @host 127.0.0.1:9050
// @BasePath /v1
func main() {
	cmd.Execute()
}
