// Application which greets you.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("hi")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet() string {
	return "Hi!"
}
