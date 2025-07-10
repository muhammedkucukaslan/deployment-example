package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Math package is ready to use.")
	http.HandleFunc("/add", AdditionHandler)
	http.ListenAndServe(":3000", nil)

}
