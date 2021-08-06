package handlers

import (
	"fmt"
	"net/http"
)

func Cal(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
	// Dưới này viết cái gì đó 
}