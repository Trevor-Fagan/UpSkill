package main

import (
	"fmt"
	"net/http"
	"log"
)

const portNumber = ":3000"

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Home")

	if err != nil {
		log.Printf("Encountered an error: %v", err)
	}
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, err := fmt.Fprintf(w, "About")
	_, _ = fmt.Fprintf(w, fmt.Sprintf("2 + 2 is %d", sum))

	if err != nil {
		log.Printf("Encountered an error: %v", err)
	}
}

// Add two values together
func AddValues(x, y int) int {
	return x + y
}

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}