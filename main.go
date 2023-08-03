package main

import (
	"fmt"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/api/prometheus/api/v1", HelloServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[5:]
	if path != "query_range" {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	} else {
		fmt.Fprint(w, "Hello World!")
	}
}
