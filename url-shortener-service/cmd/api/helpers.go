package main

import (
	"context"
	"fmt"
	pb "urlShortener/proto"
)

func (s *Server) AddUrlShortener(ctx context.Context, request *pb.UrlShortenerRequest) (*pb.SetUrlShorteneResponse, error) {

	urlPath := request.GetUrlPath()
	// short url here
	fmt.Println("the url path", urlPath)
	//return response
	return &pb.SetUrlShorteneResponse{Result: "result"}, nil

}
