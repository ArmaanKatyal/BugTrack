package main

import (
	"Backend/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/api/v1/project", controllers.AllProjects).Methods("GET")
	router.HandleFunc("/api/v1/project/create", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/project/update/{id}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/api/v1/project/delete/{id}", controllers.DeleteProject).Methods("DELETE")
	router.HandleFunc("/api/v1/ticket", controllers.AllTickets).Methods("GET")
	router.HandleFunc("/api/v1/ticket/create", controllers.CreateTicket).Methods("POST")
	router.HandleFunc("/api/v1/ticket/update/{id}", controllers.UpdateTicket).Methods("PUT")
	router.HandleFunc("/api/v1/ticket/delete/{id}", controllers.DeleteTicket).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
