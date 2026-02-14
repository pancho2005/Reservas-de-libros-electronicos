package usuarios

// Usuario representa a una persona registrada en el sistema.
type Usuario struct {
	id       int
	nombre   string
	apellido string
}

// Constructor
func NewUsuario(id int, nombre, apellido string) *Usuario {
	return &Usuario{id, nombre, apellido}
}

// Encapsulación
func (u *Usuario) ID() int {
	return u.id
}

func (u *Usuario) NombreCompleto() string {
	return u.nombre + " " + u.apellido
}
