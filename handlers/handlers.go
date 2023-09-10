package handlers

import (
	"fmt"
	"net/http"
)

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
