package main

import (
	"fmt"
	"net/http"
)

const addr = ":8000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
		fmt.Println("Endpoint hello world hit")
	})

	fmt.Printf("starting server at %s \n", addr)
	http.ListenAndServe(addr, nil)
}
