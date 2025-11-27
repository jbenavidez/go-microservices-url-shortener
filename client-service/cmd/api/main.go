package main

import (
	pb "client/proto/generated"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = 8080

type application struct {
	DSN        string
	GRPCClient pb.UrlShortenerServiceClient
}

func main() {

	fmt.Println("init client")
	//set app
	var app application
	// set gRPC connection
	conn, err := grpc.Dial("url-shortener-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	//set gRPC Client
	client := pb.NewUrlShortenerServiceClient(conn)
	app.GRPCClient = client

	log.Println("Starting application on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
