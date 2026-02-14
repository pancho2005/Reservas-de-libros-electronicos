package reportes

import "biblioteca-go/reservas"

// Función pura: no modifica estado
func TotalReservas(lista []*reservas.Reserva) int {
	return len(lista)
}
