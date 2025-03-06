package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SirishaGopigiri/testing-http-server/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/users", handlers.UserHandler)
	http.HandleFunc("/user/get", handlers.GetUserHandler)
	http.HandleFunc("/user/create", handlers.CreateUserHandler)
	http.HandleFunc("/user/update", handlers.UpdateUserHandler)
	http.HandleFunc("/user/delete", handlers.DeleteUserHandler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

	port := "8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
