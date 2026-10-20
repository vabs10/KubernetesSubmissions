package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func routeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "todo-server is running")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", routeHandler)

	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
