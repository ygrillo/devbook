package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios buscar todos os usuários no banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários."))
}

// BuscarUsuario busca um usuário salvo no banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário."))
}

// AtualizarUsuario alterar as informações de um usuário no banco de dados.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário."))
}

// DeletarUsuario exclui as informações de um usuário no banco de dados.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário."))
}
