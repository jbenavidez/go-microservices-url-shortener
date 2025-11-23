package main

import (
	"fmt"
	pb "urlShortener/proto"
)

type Server struct {
	pb.UnimplementedUrlShortenerServiceServer
}

func main() {
	fmt.Println("starting url-shortener service...")
}
