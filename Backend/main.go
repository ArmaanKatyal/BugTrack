package main

import (
	"Backend/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/project/create", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/ticket/create/{projectID}", controllers.CreateTicket).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
