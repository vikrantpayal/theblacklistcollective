package main

import (
	"fmt"
	"log"
	"net/http"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from the Go server!")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Serve static files from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/greeting", greetingHandler)
	http.HandleFunc("/api/submit", submitHandler)

	port := ":8000"
	fmt.Printf("Server starting at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
