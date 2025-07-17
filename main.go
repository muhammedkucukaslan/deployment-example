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

	mux := http.NewServeMux()
	mux.HandleFunc("/add", AdditionHandler)
	mux.HandleFunc("/env", DotenvHandler)
	mux.HandleFunc("/hello", HelloWorldHandler)

	fmt.Println("Server is running on http://localhost:" + PORT + "/")
	http.ListenAndServe(":"+PORT, LoggerMiddleware(mux))

}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func DotenvHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", os.Getenv("NAME"))
}
