package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", printHeader)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func printHeader(w http.ResponseWriter, r *http.Request) {
	for k, vs := range r.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	if r.Header.Get("X-FORWARDED-FOR") != "" {
		fmt.Println("forwarded:", r.Header.Get("X-FORWARDED-FOR"))
	}
	fmt.Println("RemoteAddr", r.RemoteAddr)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
