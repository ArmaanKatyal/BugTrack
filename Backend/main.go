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
	router.HandleFunc("/api/v1/project/{id:[0-9a-fA-F]{24}}", controllers.Project).Methods("GET")
	router.HandleFunc("/api/v1/project/create", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/project/update/{id:[0-9a-fA-F]{24}}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/api/v1/project/delete/{id:[0-9a-fA-F]{24}}", controllers.DeleteProject).Methods("DELETE")
	router.HandleFunc("/api/v1/ticket", controllers.AllTickets).Methods("GET")
	router.HandleFunc("/api/v1/ticket/{id:[0-9a-fA-F]{24}}", controllers.Ticket).Methods("GET")
	router.HandleFunc("/api/v1/ticket/create", controllers.CreateTicket).Methods("POST")
	router.HandleFunc("/api/v1/ticket/update/{id:[0-9a-fA-F]{24}}", controllers.UpdateTicket).Methods("PUT")
	router.HandleFunc("/api/v1/ticket/delete/{id:[0-9a-fA-F]{24}}", controllers.DeleteTicket).Methods("DELETE")
	router.HandleFunc("/api/v1/user", controllers.AllUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.User).Methods("GET")
	router.HandleFunc("/api/v1/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/user/validUsername/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.CheckUsernameExists).Methods("POST")
	router.HandleFunc("/api/v1/user/profile/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.UserProfile).Methods("GET")
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
