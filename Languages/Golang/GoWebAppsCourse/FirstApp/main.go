package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":3000"

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}