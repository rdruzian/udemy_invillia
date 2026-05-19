package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type Usuario struct {
	Nome  string
	Email string
}

func main() {
	u := Usuario{
		Nome:  "Teste",
		Email: "abcd@gmail.com",
	}

	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	//Cria o servidor e configura a porta
	log.Fatal(http.ListenAndServe(":5000", nil))
}
