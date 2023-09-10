package main

import (
	"fmt"
	"net/http"
)

const port = ":8084"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Welcome to home page")

	if err != nil {
		fmt.Println("Error in home page")
		fmt.Println(err)
	}

	fmt.Println("Number of bytes send from Home:", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is About page")

	if err != nil {
		fmt.Println("Error in about page")
		fmt.Println(err)
	}

	fmt.Println("Number of bytes send from about:", n)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server running on port", port)
	_ = http.ListenAndServe(port, nil)
}
