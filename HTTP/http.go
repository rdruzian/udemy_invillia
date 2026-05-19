package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func main() {
	//                 URI
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	//Cria o servidor e configura a porta
	log.Fatal(http.ListenAndServe(":5000", nil))
}
