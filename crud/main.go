package main

import (
	"crud/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", server.BuscaUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.BuscaUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", server.DeletarUsuario).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))

}
