package main

import (
	"fmt"
	"net/http"

	"github.com/jmohandas/usermanager/handlers"
)

const port = ":8084"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server running on port", port)
	_ = http.ListenAndServe(port, nil)
}
