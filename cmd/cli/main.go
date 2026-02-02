package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Get hello world on localhost:8080
	resp, err := http.Get("http://localhost:8080?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
