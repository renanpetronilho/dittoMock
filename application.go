package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CallData struct {
	Method     string
	Timeout    int
	StatusCode int
	Headers    []string
}

func handler(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	timeout, _ := strconv.Atoi(r.Header.Get("timeout"))
	statusCode, _ := strconv.Atoi(r.Header.Get("statusCode"))

	time.Sleep(time.Duration(timeout) * time.Second)

	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}

	calldata := CallData{Method: method, Timeout: timeout, StatusCode: statusCode, Headers: keys}

	js, err := json.Marshal(calldata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
