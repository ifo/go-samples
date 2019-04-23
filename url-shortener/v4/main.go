package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	if r.URL.Path != "/" {
		idStr := r.URL.Path[1:]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "I don't know this short url", http.StatusBadRequest)
			return
		}
		if redirectURL, exists := shortURLs[id]; !exists {
			http.Error(w, "I don't know this short url", http.StatusBadRequest)
			return
		} else {
			http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
			return
		}
	}

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
