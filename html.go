package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

const tmpl = ""

func badHTMLTemplate() {
	a := "something from another place"
	t := template.Must(template.New("ex").Parse(tmpl))
	v := map[string]interface{}{
		"Title": "Test <b>World</b>",
		"Body":  template.HTML(a),
	}
	t.Execute(os.Stdout, v)
}

// XSS: directly writing user input to response
func xssHandler(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("q")
	fmt.Fprintf(w, "<html><body>Search results for: %s</body></html>", userInput)
}

// Open redirect vulnerability
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	http.Redirect(w, r, target, http.StatusFound)
}

// SSRF: fetching user-controlled URL
func fetchURL(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("fetch")
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "Status: %d", resp.StatusCode)
}
