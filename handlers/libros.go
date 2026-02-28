package handlers

import (
	"biblioteca-go/data"
	"biblioteca-go/models"
	"encoding/json"
	"net/http"
)

func CrearLibro(w http.ResponseWriter, r *http.Request) {

	data.Mutex.Lock()
	defer data.Mutex.Unlock()

	var libro models.Libro
	json.NewDecoder(r.Body).Decode(&libro)

	libro.ID = len(data.Libros) + 1
	data.Libros = append(data.Libros, libro)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libro)
}

func ListarLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Libros)
}
