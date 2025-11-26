package main

import (
	"client/cmd/api/models"
	pb "client/proto/generated"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateUrlShortener
func (app *application) CreateUrlShortener(w http.ResponseWriter, r *http.Request) {
	// Get payload
	var payload models.UrlShortener

	err := app.readJSON(w, r, &payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := grpc.Dial("url-shortener-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := pb.NewUrlShortenerServiceClient(conn)

	//set request
	req := &pb.AddUrlShortenerRequest{
		UrlPath: payload.FullPath,
	}
	//call grpc
	response, err := client.AddUrlShortener(r.Context(), req)
	if err != nil {
		fmt.Println("something break", err)
		return
	}

	payload.Shortcut = response.Result

	resp := JSONResponse{
		Error:   false,
		Message: "URL Shortener successfully",
		Data:    payload,
	}
	// send resposne
	_ = app.writeJSON(w, http.StatusOK, resp)

}

// GetAllUrlShorteners get all recods
func (app *application) GetAllUrlShorteners(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial("url-shortener-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := pb.NewUrlShortenerServiceClient(conn)
	response, err := client.GetAllUrlShorteners(r.Context(), &emptypb.Empty{})
	if err != nil {
		fmt.Println("something break", err)
		fmt.Fprint(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "retrieved URL Shortener successfully",
		Data:    response,
	}

	_ = app.writeJSON(w, http.StatusOK, resp)

}

// UpdateUrlShortener: update record
func (app *application) UpdateUrlShortener(w http.ResponseWriter, r *http.Request) {
	urlID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// get payload
	var payload pb.UrlShortener
	err = app.readJSON(w, r, &payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	payload.Id = int64(urlID)
	conn, err := grpc.Dial("url-shortener-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := pb.NewUrlShortenerServiceClient(conn)
	//set request
	req := &pb.UpdateUrlShortenerRequest{
		Payload: &payload,
	}
	//call grpc
	newShorcut, err := client.UpdateUrlShortener(r.Context(), req)
	if err != nil {
		fmt.Println("something break", err)
		return
	}
	payload.Shortcut = newShorcut.Result
	resp := JSONResponse{
		Error:   false,
		Message: " URL was updated successfully",
		Data:    &payload,
	}

	_ = app.writeJSON(w, http.StatusOK, resp)

}
