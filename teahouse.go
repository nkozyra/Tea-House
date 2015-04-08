package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

const (
	PORT = ":8080"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type Page struct {
	Title string
	Body  string
}

func loadPage(title string) Page {
	return Page{Title: title, Body: ""}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	p := loadPage("login")
	// t, _ := template.ParseFiles("templates/login.html")
	templates.Execute(w, p)
}

func main() {
	fmt.Println("Starting")
	r := mux.NewRouter()

	r.HandleFunc("/login", LoginHandler)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}
