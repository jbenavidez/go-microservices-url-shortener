package main

import (
	"context"
	"fmt"
	"net"

	"urlShortener/models"
	pb "urlShortener/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening gRPC......")
	srv := grpc.NewServer()
	pb.RegisterUrlShortenerServiceServer(srv, &server{})

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) AddUrlShortener(ctx context.Context, request *pb.AddUrlShortenerRequest) (*pb.AddUrlShorteneResponse, error) {

	urlPath := request.GetUrlPath()

	//

	//
	var urlShortener models.UrlShortener
	urlShortener.FullPath = urlPath
	urlShortener.Shortcut = "in dev"
	//save
	_, err := app.DB.CreateUrlShortener(urlShortener)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	//return response
	return &pb.AddUrlShorteneResponse{Result: urlShortener.Shortcut}, nil

}
