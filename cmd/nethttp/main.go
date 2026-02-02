package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (

	// Handled formats
	JSON      = "json"
	XML       = "xml"
	Unhandled = "unhandled"

	// Query parameters
	format = "format"

	// Content types
	applicationJSON = "application/json"
	applicationXML  = "application/xml"

	// Default message
	defaultMessage = "Hello World"
)

type HelloWorld struct {
	Message string `json:"msg" xml:"content"`
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		hello := HelloWorld{Message: defaultMessage}

		switch r.URL.Query().Get(format) {
		case JSON:
			w.Header().Set("Content-Type", applicationJSON)
			content, _ := json.Marshal(hello)
			w.Write(content)
		case XML:
			w.Header().Set("Content-Type", applicationXML)
			content, _ := xml.Marshal(hello)
			w.Write(content)
		default:
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("Format not handled"))
		}
	})

	http.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
		hello := HelloWorld{Message: defaultMessage + r.PostFormValue("name")}

		switch r.URL.Query().Get(format) {
		case JSON:
			w.Header().Set("Content-Type", applicationJSON)
			content, err := json.Marshal(hello)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Internal server error"))
				return
			}
			w.Write(content)
		case XML:
			w.Header().Set("Content-Type", applicationXML)
			content, err := xml.Marshal(hello)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Internal server error"))
				return
			}
			w.Write(content)
		default:
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("Format not handled"))
		}

	})

	http.ListenAndServe(":8080", nil)
}
