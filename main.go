package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Simple route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! I'm Kalyan, a Golang & Python Developer based in Bangalore.")
	})

	fmt.Println("Server is starting on port 8080...")
	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
