package main

import (
	"html/template"
	"log"
	"net/http"
)

// Create global template pointer.
var tmpl *template.Template

type PageData struct {
	Title string
}

func main() {
	// Create the mux to handle HTTP requests.
	mux := http.NewServeMux()
	mux.HandleFunc("/about/", aboutHandler)
	mux.HandleFunc("/contact/", contactHandler)
	mux.HandleFunc("/", indexHandler)

	// Create fileserver to handle incoming file requests.
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Run local webserver.
	log.Println("Running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// indexHandler is the template handler for the home page (static/html/index.html).
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("static/html/index.html"))
	var data = PageData{Title: "Home"}

	// Execute template.
	tmpl.Execute(w, data)
}

// aboutHandler is the template handler for the about page (static/html/about.html).
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("static/html/about.html"))
	var data = PageData{Title: "About"}

	// Execute template.
	tmpl.Execute(w, data)
}

// contactHandler is the template handler for the contact page (static/html/contact.html).
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("static/html/contact.html"))
	var data = PageData{Title: "Contact"}

	// Execute template.
	tmpl.Execute(w, data)
}
