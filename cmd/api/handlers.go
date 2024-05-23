package main

import (
	"fmt"
	"net/http"
)

// Hello is the exported function that handles incoming requests
// and it will be used as a router parameter in http.HandleFunc
// ***
// by adding app *application, same as routes func
// hello func has access to all the info stored in application
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go server! %s", app.Domain)
}
