// simple_server_01.go
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe("localhost:8080", nil)
}
