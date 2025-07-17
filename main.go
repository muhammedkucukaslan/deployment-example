package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	fmt.Println("Math package is ready to use.")
	http.HandleFunc("/add", AdditionHandler)
	http.HandleFunc("/env", DotenvHandler)
	http.HandleFunc("/hello", HelloWorldHandler)
	http.HandleFunc("/square", SquareHandler)
	fmt.Println("Server is running on http://localhost:" + PORT + "/")
	http.ListenAndServe(":"+PORT, nil)

}

func DotenvHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", os.Getenv("NAME"))
}
