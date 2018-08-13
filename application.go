package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	timeout, _ := strconv.Atoi(r.Header.Get("timeout"))
	statusCode, _ := strconv.Atoi(r.Header.Get("statusCode"))

	fmt.Println("Method -> ", method, " Timeout -> ", timeout, "statusCode ->", statusCode)

	time.Sleep(time.Duration(timeout) * time.Second)

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Path agora %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
