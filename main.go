package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", printHeader)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func printHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println(k, v)
	}
}
