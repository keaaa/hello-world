package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addr := os.Getenv("URL_SERVER")
		if addr == "" {
			addr = "http://:8000"
		}

		resp, err := http.Get(addr)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		fmt.Fprintf(w, "<h1>%s</h1>", string(body))
	})

	http.ListenAndServe(":8080", nil)
}
