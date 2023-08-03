package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var port = os.Getenv("PORT")

type dataStruct struct {
	Status string      `json:"status"`
	Data   queryResult `json:"data"`
}

type ValueType int
type Value interface {
	Type() ValueType
	String() string
}

type queryResult struct {
	Type   string      `json:"resultType"`
	Result interface{} `json:"result"`
}

func main() {
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/api/v1/query", QueryServer)
	http.HandleFunc("/api/v1/query_range", QueryRangeServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func QueryServer(w http.ResponseWriter, r *http.Request) {
	res := queryResult{Type: "vector", Result: []string{}}
	data := dataStruct{Status: "success", Data: res}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func QueryRangeServer(w http.ResponseWriter, r *http.Request) {
	res := queryResult{Type: "vector", Result: []string{}}
	data := dataStruct{Status: "success", Data: res}
	w.Header().Set("Content-Type", "application/json")
	time.Sleep(11 * time.Second)
	json.NewEncoder(w).Encode(data)
}
