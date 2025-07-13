package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Math package is ready to use.")
	http.HandleFunc("/add", AdditionHandler)
	http.HandleFunc("/hello", HelloWorldHandler)
	http.HandleFunc("/square", SquareHandler)
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe("0.0.0.0:3000", nil)

}
