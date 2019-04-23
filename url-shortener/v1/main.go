package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server starting on port 3030")
	err := http.ListenAndServe(":3030", nil)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the index page!")
}
