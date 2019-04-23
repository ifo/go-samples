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
	page := `
<html><body>
Enter the URL you want to shorten!
<form method="post" action="/new">
	<input type="text" name="url" placeholder="example.com" />
	<button type="submit">Shorten!</button>
</form>
</body></html>
`

	fmt.Fprint(w, page)
}
