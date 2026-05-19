package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var usuario usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		fmt.Sprintf("erro to convert byte in struct %v", err)
		w.Write([]byte("Erro ao converter em usuario"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		fmt.Sprintf("erro to connect to database %v", err)
		w.Write([]byte("Erro ao conectar no banco"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if err != nil {
		fmt.Sprintf("erro to prepare statement %v", err)
		w.Write([]byte("Erro ao criar statement"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	insert, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		fmt.Sprintf("erro to insert user %v", err)
		w.Write([]byte("Erro ao inserir usuário"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := insert.LastInsertId()
	if err != nil {
		fmt.Sprintf("erro to get last inserir id %v", err)
		w.Write([]byte("Erro ao buscar id do usuário"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário adicionado com id %d", id)))
}

func BuscaUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao ler o usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var usuario usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuarios"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter usuários para json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter id para inteiro"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao ler o usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM usuarios where id = ?", uint32(ID))
	if err != nil {
		w.Write([]byte("Erro ao buscar usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer row.Close()

	var usuario usuario
	if row.Next() {
		if err := row.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuarios"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao converter usuários para json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario usuario

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter id para inteiro"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &usuario); err != nil {
		fmt.Sprintf("erro to convert byte in struct %v", err)
		w.Write([]byte("Erro ao converter em usuario"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao ler o usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar usuarios"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter id para inteiro"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao ler o usuario"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("delete FROM usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar usuario"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
