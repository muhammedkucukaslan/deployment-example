package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Add(a, b int) int {
	return a + b
}

func AdditionHandler(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)
	result := Add(aInt, bInt)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"result": %d}`, result)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}
