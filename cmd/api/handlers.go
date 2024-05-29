package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nstutorial-server/internal/models"
)

// Home is the exported function that handles incoming requests
// and it will be used as a router parameter in http.HandleFunc
// ***
// by adding app *application, same as routes func
// Home func has access to all the info stored in application
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
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

func (app *application) AllLessons(w http.ResponseWriter, r *http.Request) {
	var lessons []models.Lesson

	lesson1 := models.Lesson{
		ID:     1,
		Level:  "Beginner",
		Lesson: 1,
		Content: models.Content{
			Title:        "Variables from BE",
			Introduction: "Variables are used to store data values in programming.",
			Examples:     "var name = 'John':",
		},
	}

	lessons = append(lessons, lesson1)

	out, err := json.Marshal(lessons)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
