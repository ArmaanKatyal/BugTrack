package main

import (
	"Backend/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func main() {

	router := mux.NewRouter() // create a new router
	router.StrictSlash(true)  // enable strict routing
	// Handle the routes
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
	router.HandleFunc("/api/v1/ticket/project/{projectID:[0-9a-fA-F]{24}}", controllers.ProjectTickets).Methods("GET")
	router.Path("/api/v1/ticket/filter/status").Queries("type", "{type}").HandlerFunc(controllers.FilterTicketsByStatus).Methods("GET")
	router.Path("/api/v1/ticket/filter/priority").Queries("type", "{type}").HandlerFunc(controllers.FilterTicketsByPriority).Methods("GET")
	router.Path("/api/v1/user").Queries("role", "{role}").HandlerFunc(controllers.AllUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.User).Methods("GET")
	router.HandleFunc("/api/v1/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/user/validUsername/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.CheckUsernameExists).Methods("GET")
	router.HandleFunc("/api/v1/user/profile/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.UserProfile).Methods("GET")
	router.HandleFunc("/api/v1/user/lock/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.Lock).Methods("GET")
	router.HandleFunc("/api/v1/user/unlock/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.UnLock).Methods("GET")
	router.HandleFunc("/api/v1/user/locked", controllers.LockedUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/role", controllers.GetUserRole).Methods("GET")
	router.HandleFunc("/api/v1/auth/login", controllers.UserLogin).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/auth/logout", controllers.UserLogout).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/auth/changePassword", controllers.ChangePassword).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/auth/forgotPassword", controllers.ForgotPassword).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/auth/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/api/v1/logs", controllers.AllLogs).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(NotFound) // set the 404 handler

	corsWrapper := cors.New(cors.Options{ // create a new cors wrapper
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "token", "Set-Cookie"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Handler:      corsWrapper.Handler(router),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	//err := http.ListenAndServe(":8080", corsWrapper.Handler(router)) // start the server
	if err != nil {
		panic(err)
	}

}

// NotFound is a 404 handler
func NotFound(w http.ResponseWriter, _ *http.Request) { // 404 handler
	w.WriteHeader(http.StatusNotFound)
}
