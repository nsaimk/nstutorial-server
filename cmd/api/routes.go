package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// 'app *application' means that any time I have a variable
// of the type application, I have access to 'routes' function
func (app *application) routes() http.Handler{
	// create a router multiplexer
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	// get request, goes to the Home handler
	mux.Get("/", app.Home)

	mux.Get("/lessons", app.AllLessons)

	return mux
}