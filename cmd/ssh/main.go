package main

import (
	"fmt"
	"github.com/iAbbos/go-grpc_auth-sso/internal/config"
)

func main() {
	//Implementing config
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//TODO: Implement logger

	//TODO: Implement application

	//TODO: Implement grpc server
}
