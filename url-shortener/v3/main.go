package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)

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
%s
</body></html>
`

	urltext := ""
	for id, sURL := range shortURLs {
		urltext += fmt.Sprintf("%d - %s<br />", id, sURL)
	}

	fmt.Fprintf(w, page, urltext)
}

var URLID = 0
var shortURLs = map[int]string{}

func newHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post requests only please!", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad form!", http.StatusBadRequest)
		return
	}

	shortURL := r.PostForm.Get("url")

	if !strings.HasPrefix(shortURL, "http://") && !strings.HasPrefix(shortURL, "https://") {
		shortURL = "http://" + shortURL
	}

	URLID++
	shortURLs[URLID] = shortURL

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
