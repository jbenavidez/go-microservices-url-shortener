package main

import (
	"net/http"
)

// CreateUser create user record
func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {

	resp := JSONResponse{
		Error:   false,
		Message: "",
	}
	_ = app.writeJSON(w, http.StatusOK, resp)

}
