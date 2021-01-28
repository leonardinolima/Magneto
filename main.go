package main

import (
	"github.com/magneto/service"
)

func main() {
	server := new(service.Server)
	server.RunServer()
}
