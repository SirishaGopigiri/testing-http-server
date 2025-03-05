package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler responds with a greeting
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
