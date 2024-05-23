package main

import (
	"fmt"
	"net/http"
)

// Function Hello is a handle function that 
//will be used as a router paramater in http.HandleFunc
func Hello2(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello go server!")
}