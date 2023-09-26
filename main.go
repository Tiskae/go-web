package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

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

	// r := mux.NewRouter()

	// File serving
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		fmt.Printf("You've requested for the book %s on page %s\n", vars["title"], vars["page"])
	})

	// // Restrict request handler to specific https method
	// r.HandleFunc("/post-result", PostResultHandler).Methods("POST")
	// r.HandleFunc("/edit-profile", EditProfileHandler).Methods("PUT")
	// r.HandleFunc("/delete-profile", DeleteProfileHandler).Methods("DELETE")

	// // Restrict request handler to specific hostnames or subdomains
	// r.HandleFunc("/make-huge-payment", HugePaymentHandler).Host("www.bankdomain.com")

	// // Restrict request handler to http/https
	// r.HandleFunc("/secure", SecureHandler).Schemes("https")
	// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// // Restrict request handler to specific prefix path
	// geoBooksRouter := r.PathPrefix("/geology-books").Subrouter()
	// geoBooksRouter.HandleFunc("/", AllBooks)
	// geoBooksRouter.HandleFunc("/{title}", GetBook)

	// COnfigure the database connection (always check for error)
	// db, err := sql.Open("mysql", "")

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	// err := db.Ping()

	http.ListenAndServe(":80", nil)
}
