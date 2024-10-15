package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func(app *application) routes() *mux.Router{
	router := mux.NewRouter()

	router.Use(secureHeaders)
	router.Use(app.logRequest)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	router.HandleFunc("/todo", app.allTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", app.oneTodo).Methods("GET")
	router.HandleFunc("/todo", app.homeCreate).Methods("POST")
	router.HandleFunc("/todo/{id}", app.homeDelete).Methods("DELETE")

	return router
}
