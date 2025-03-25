package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello, HTTP! This is the response\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handler has been initialized")
	fmt.Printf("Listening on localhost:8080\n")
	io.WriteString(w, "Hello, HTTP!\n")

}



func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")

	} else if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
