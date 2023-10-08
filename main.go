package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/flag", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server")
	})

	port := ":4372" // Порт, на котором будет слушать сервер
	http.ListenAndServe(port, nil)
}
