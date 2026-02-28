package handlers

import (
	"biblioteca-go/data"
	"encoding/json"
	"net/http"
)

func TotalReservas(w http.ResponseWriter, r *http.Request) {

	resultado := map[string]int{
		"total_reservas": len(data.Reservas),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}

func LibrosDisponibles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Libros)
}
