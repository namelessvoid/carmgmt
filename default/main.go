package main

import (
	"os"
	"log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
}

func main() {
    http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
	}

    log.Fatal(http.ListenAndServe(":"+port, nil))
}