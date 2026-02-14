package libros

import "biblioteca-go/errores"

// Libro representa un libro electrónico del sistema.
type Libro struct {
	titulo string
	autor  string
	stock  int
}

// Constructor
func NewLibro(titulo, autor string, stock int) *Libro {
	return &Libro{titulo, autor, stock}
}

func (l *Libro) Titulo() string {
	return l.titulo
}

func (l *Libro) Stock() int {
	return l.stock
}

// Reduce el stock validando disponibilidad
func (l *Libro) ReducirStock() error {
	if l.stock <= 0 {
		return errores.ErrStockInsuficiente
	}
	l.stock--
	return nil
}
