package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Home is the exported function that handles incoming requests
// and it will be used as a router parameter in http.HandleFunc
// ***
// by adding app *application, same as routes func
// Home func has access to all the info stored in application
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status: "active",
		Message: "Learn with NS Tutorial",
		Version: "0.1.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
