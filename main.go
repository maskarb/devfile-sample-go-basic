package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/api/prometheus/api/v1/query", QueryServer)
	http.HandleFunc("/api/prometheus/api/v1/query_range", QueryRangeServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func QueryServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Query World!")
}

func QueryRangeServer(w http.ResponseWriter, r *http.Request) {
	time.Sleep(11 * time.Second)
	fmt.Fprint(w, "Hello QueryRange World!")
}
