package main

import (
	"fmt"
	"log"
	"net/http"
)

func page(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println(r)
	id := r.URL.Query().Get("id")
	fmt.Println("id", id)

}

func startWebServer() {

	http.HandleFunc("/", page)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":7777", nil))
}
