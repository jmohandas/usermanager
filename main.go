package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world in go!")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Bytes send:", n)
	})

	_ = http.ListenAndServe(":8084", nil)
}
