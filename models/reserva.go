package models

type Reserva struct {
	ID        int `json:"id"`
	UsuarioID int `json:"usuario_id"`
	LibroID   int `json:"libro_id"`
}
