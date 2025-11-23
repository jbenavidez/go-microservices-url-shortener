package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN string
}

func main() {

	fmt.Println("init client")
	//set app
	var app application

	log.Println("Starting application on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
