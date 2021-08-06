package server

import (
	"fmt"
	"net/http"

	"../handlers"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8090/cal")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/cal", handlers.Cal) // simple hello
	
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic("Error when running server")
	}
}