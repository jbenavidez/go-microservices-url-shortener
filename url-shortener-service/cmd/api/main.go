package main

import (
	"fmt"
	"log"
	"urlShortener/repository"
	dbrepo "urlShortener/repository/db_repo"
)

const (
	gRpcPort = "50001"
)

type Config struct {
	DSN string
	DB  repository.DatabaseRepo
}

func main() {
	fmt.Println("starting url-shortener service...")
	app := Config{}

	conn := app.connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	//set up helper
	NewGrpcHelper(&app)
	// Set up gRPC
	app.gRPCListenAndServe()
}
