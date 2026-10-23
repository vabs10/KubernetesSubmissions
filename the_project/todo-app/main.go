package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func routeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `
 <html>
		<head>
			<title>Kubernetes Submission</title>
		</head>
		<body style="text-align:center; margin-top:200px;">
		<h1>Welcome to KubernetesSubmission!</h1>
		<p>Your application is runnning successfully on Kubernetes.</p>
		</body>
		</html>
		`)
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
