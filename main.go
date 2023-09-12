package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Dynamic routes
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Hello, welcome to the homepage")
	// })

	// http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Here's my blog page")
	// })

	// http.HandleFunc("/blog/bmw-m4-f82", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("The M4 F82 is a German masterpiece from BMW. It boasts a powerful V8 engine producing an astonishing 512 HP stock. It ...")
	// })

	// // File serving
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		fmt.Printf("You've requested for the book %s on page %s\n", vars["title"], vars["page"])
	})

	http.ListenAndServe(":80", r)
}
