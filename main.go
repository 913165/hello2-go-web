package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Greetings,Hello Golang Web.\n")
}

func main() {

	http.HandleFunc("/", hello)

	http.ListenAndServe(":3000", nil)
}
