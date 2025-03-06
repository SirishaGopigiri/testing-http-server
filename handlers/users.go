package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct represents user data
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var users = map[string]User{}

// UserHandler returns user data as JSON
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	users[newUser.Name] = newUser
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if _, exists := users[name]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	users[updatedUser.Name] = updatedUser
	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if _, exists := users[name]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	delete(users, name)
	w.WriteHeader(http.StatusNoContent)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if _, exists := users[name]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	user := users[name]
	json.NewEncoder(w).Encode(user)
}
