package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type pessoas struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	p1 := pessoas{"Pérola", "diego@gmail.com"}
	templates.ExecuteTemplate(w, "home.html", p1)
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	p2 := pessoas{"Maria", "maria@gmail.com"}
	templates.ExecuteTemplate(w, "usuarios.html", p2)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	fmt.Println("Escutando na porta 5010...")
	log.Fatal(http.ListenAndServe(":5010", nil))

}
