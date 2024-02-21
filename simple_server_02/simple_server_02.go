// simple_server_02.go
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintln(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe("localhost:8080", nil)
}
