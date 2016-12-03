package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	decodedString, err := base64.StdEncoding.DecodeString(r.URL.Path[1:])
	if err != nil {
		w.Write([]byte("Not a valid base64 string"))
		return
	}

	w.Write(decodedString)
}

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	decodedString, err := base64.StdEncoding.DecodeString(r.URL.Path[1:])
	if err != nil {
		w.Write([]byte("Not a valid base64 string"))
		return
	}

	pageData := "<script>" + string(decodedString) + "</script>"
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
