package handlers

import (
	"fmt"
	"net/http"
)

func GetResponse() string {
	return "Hello world!!"
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, GetResponse())
	fmt.Println("Endpoint hello world hit")
}
