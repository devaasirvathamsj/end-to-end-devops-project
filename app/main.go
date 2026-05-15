package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	appName := os.Getenv("APP_NAME")
	appVersion := os.Getenv("APP_VERSION")

	// Serve HTML
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "App: %s, Version: %s", appName, appVersion)
	})

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
