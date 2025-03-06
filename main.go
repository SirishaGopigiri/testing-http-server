package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SirishaGopigiri/testing-http-server/handlers"
)

func main() {

	r := mux.NewRouter()

	// Middleware
	r.Use(handlers.LoggingMiddleware)

	// Routes with method-based handling
	r.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.UsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", handlers.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handlers.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handlers.UpdateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", handlers.DeleteUserHandler).Methods(http.MethodDelete)

	port := "8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, r))
}
