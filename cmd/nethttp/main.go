package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World " + r.PostFormValue("name")))
	})

	http.ListenAndServe(":8080", nil)
}
