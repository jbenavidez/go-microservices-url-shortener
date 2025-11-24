package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"urlShortener/models"
	pb "urlShortener/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const shortenerLenght = 6

var app *Config

// NewGrpcHelper make app config avaible
func NewGrpcHelper(a *Config) {
	app = a
}

type server struct {
	pb.UnimplementedUrlShortenerServiceServer
}

// gRPCListenAndServe set up gRPC conenction
func (app *Config) gRPCListenAndServe() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterUrlShortenerServiceServer(srv, &server{})

	reflection.Register(srv)
	log.Printf("gRPC server started on port %s ", gRpcPort)
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) AddUrlShortener(ctx context.Context, request *pb.AddUrlShortenerRequest) (*pb.AddUrlShorteneResponse, error) {

	urlPath := request.GetUrlPath()

	//
	shortener := app.GenerateUniqueStringFromLongUrlPath(urlPath, shortenerLenght)
	//
	var urlShortener models.UrlShortener
	urlShortener.FullPath = urlPath
	urlShortener.Shortcut = shortener
	//save
	// _, err := app.DB.CreateUrlShortener(urlShortener)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return nil, err
	// }
	// //return response
	return &pb.AddUrlShorteneResponse{Result: urlShortener.Shortcut}, nil

}
