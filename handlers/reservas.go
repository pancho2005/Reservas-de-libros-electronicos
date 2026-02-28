package handlers

import (
	"biblioteca-go/data"
	"biblioteca-go/models"
	"encoding/json"
	"net/http"
)

func CrearReserva(w http.ResponseWriter, r *http.Request) {

	data.Mutex.Lock()
	defer data.Mutex.Unlock()

	var reserva models.Reserva
	json.NewDecoder(r.Body).Decode(&reserva)

	// Buscar libro
	for i := range data.Libros {
		if data.Libros[i].ID == reserva.LibroID {

			// Intentar reducir stock usando método
			err := data.Libros[i].ReducirStock()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			reserva.ID = len(data.Reservas) + 1
			data.Reservas = append(data.Reservas, reserva)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reserva)
			return
		}
	}

	http.Error(w, "Libro no encontrado", http.StatusNotFound)
}

func ListarReservas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Reservas)
}
