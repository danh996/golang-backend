package main

import (
	"fmt"
	"io"
	"net/http"
)

const webPort = ":8008"

func main() {
	fmt.Println("School App")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(webPort, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("HTTP Server is running on Port: ", webPort)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
