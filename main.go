package main

import (
	"fmt"
	"net/http"

	"github.com/keaaa/hello-world/handlers"
)

const addr = ":8000"

func main() {
	http.HandleFunc("/", handlers.HelloWorld)

	fmt.Printf("starting server at %s \n", addr)
	http.ListenAndServe(addr, nil)
}
