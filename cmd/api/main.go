package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

// Hello is the exported function that handles incoming requests
//and it will be used as a router paramater in http.HandleFunc
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello go server!")
}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to the database
	app.Domain = "example.com"

	log.Println("Starting the app on port 8080")

	http.HandleFunc("/", Hello)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
