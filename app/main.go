package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	appName := os.Getenv("APP_NAME")
	appVersion := os.Getenv("APP_VERSION")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "App: %s, Version: %s", appName, appVersion)
	})

	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
