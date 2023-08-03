package main

import (
	"encoding/json"
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
	http.HandleFunc("/api/v1/query", QueryServer)
	http.HandleFunc("/api/v1/query_range", QueryRangeServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func QueryServer(w http.ResponseWriter, r *http.Request) {
	data := struct{ Status string }{Status: "query success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func QueryRangeServer(w http.ResponseWriter, r *http.Request) {
	data := struct{ Status string }{Status: "query range success"}
	w.Header().Set("Content-Type", "application/json")
	time.Sleep(11 * time.Second)
	json.NewEncoder(w).Encode(data)
}
