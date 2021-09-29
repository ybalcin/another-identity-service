package main

import (
	"github.com/ybalcin/another-identity-service/server"
	_ "github.com/ybalcin/another-identity-service/startup"
)

func main() {
	server.Serve()
}
