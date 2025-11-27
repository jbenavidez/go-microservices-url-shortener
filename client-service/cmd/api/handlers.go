package main

import (
	"client/cmd/api/models"
	pb "client/proto/generated"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	//set request
	req := &pb.AddUrlShortenerRequest{
		UrlPath: payload.FullPath,
	}
	//call grpc
	response, err := app.GRPCClient.AddUrlShortener(r.Context(), req)
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

	response, err := app.GRPCClient.GetAllUrlShorteners(r.Context(), &emptypb.Empty{})
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

	//set request
	req := &pb.UpdateUrlShortenerRequest{
		Payload: &payload,
	}
	//call grpc
	newShorcut, err := app.GRPCClient.UpdateUrlShortener(r.Context(), req)
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

// GetUrlShortener get url
func (app *application) GetUrlShortener(w http.ResponseWriter, r *http.Request) {
	shortcut := chi.URLParam(r, "shortcut")
	if len(shortcut) == 0 {
		app.errorJSON(w, errors.New("invalid shortcut"))
		return
	}
	fmt.Println("the value ", shortcut)

	//set request
	req := &pb.GetUrlShortenerRequest{
		Shortcut: shortcut,
	}

	response, err := app.GRPCClient.GetUrlShortener(r.Context(), req)
	if err != nil {
		fmt.Println("something break", err)
		return
	}
	resp := JSONResponse{
		Error:   false,
		Message: " URL was updated successfully",
		Data:    response.Result,
	}

	_ = app.writeJSON(w, http.StatusOK, resp)

}
