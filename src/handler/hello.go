package handler

import (
	"fmt"
	"net/http"
)

// Hello is handler
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
