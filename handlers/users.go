package handlers

import (
	"encoding/json"
	"net/http"
)

// User struct represents user data
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// UserHandler returns user data as JSON
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Sample user data
	user := User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
