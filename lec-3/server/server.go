package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trandtung1209/calculator/calc"
)

func RunServer() {
	
	fmt.Println("Server is listening in http://localhost:8080")
	defer func() {
		fmt.Println("Server is Stopped")
	}()

	http.HandleFunc("/submit", calc.Submit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}