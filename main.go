package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write “Hello World” to the response body
	fmt.Fprint(w, "Hello World")
}

func main() {
	// Register the handler for the root path
	http.HandleFunc("/", helloHandler)

	// Start the server on port 80; log.Fatal will print any error and exit
	log.Println("Starting server on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
