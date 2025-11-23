package main

import (
	"fmt"
	pb "urlShortener/proto"
)

type server struct {
	pb.UnimplementedUrlShortenerServiceServer
}

func main() {
	fmt.Println("starting url-shortener service...")
}
