package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Dynamic routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, welcome to the homepage")
	})

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Here's my blog page")
	})

	http.HandleFunc("/blog/bmw-m4-f82", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("The M4 F82 is a German masterpiece from BMW. It boasts a powerful V8 engine producing an astonishing 512 HP stock. It ...")
	})

	// File serving
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
