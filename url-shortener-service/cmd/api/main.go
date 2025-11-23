package main

import (
	"fmt"
	"urlShortener/repository"
)

const (
	rpcPort = "5001"
)

type Config struct {
	DB repository.DatabaseRepo
}

func main() {
	fmt.Println("starting url-shortener service...")
	app := Config{}
	NewGrpcHelper(&app)
	// Set up gRPC
	app.gRPCListenAndServe()
}
