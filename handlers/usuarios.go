package handlers

import (
	"biblioteca-go/data"
	"biblioteca-go/models"
	"encoding/json"
	"net/http"
)

func CrearUsuario(w http.ResponseWriter, r *http.Request) {

	data.Mutex.Lock()
	defer data.Mutex.Unlock()

	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	usuario.ID = len(data.Usuarios) + 1
	data.Usuarios = append(data.Usuarios, usuario)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Usuarios)
}
