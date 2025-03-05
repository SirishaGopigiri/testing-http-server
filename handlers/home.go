package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler serves the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!")
}
