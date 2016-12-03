package main

import (
	"fmt"
	"net/http"
	"os"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	pageData := r.URL.Path[1:]
	w.Write([]byte(pageData))
}

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	pageData := "<script>" + r.URL.Path[1:] + "</script>"
	w.Write([]byte(pageData))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/file/", fileHandler)
	http.HandleFunc("/script/", scriptHandler)
	http.HandleFunc("/", scriptHandler) // handle scripts by default

	fmt.Println("Running on port", port)
	http.ListenAndServe(":"+port, nil)
}
