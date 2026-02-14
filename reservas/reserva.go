package reservas

import (
	"biblioteca-go/libros"
	"biblioteca-go/usuarios"
)

// Reserva relaciona un usuario con un libro.
type Reserva struct {
	Usuario *usuarios.Usuario
	Libro   *libros.Libro
}
