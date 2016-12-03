package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := "<script>" + r.URL.Path[1:] + "</script>"
		w.Write([]byte(pageData))
	})

	fmt.Println("Running on port", port)
	http.ListenAndServe(":"+port, nil)
}
