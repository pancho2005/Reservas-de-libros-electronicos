package reservas

// GestionReservas administra las reservas realizadas.
type GestionReservas struct {
	lista []*Reserva
}

// Constructor
func NewGestionReservas() *GestionReservas {
	return &GestionReservas{}
}

// RegistrarReserva valida el stock y registra la reserva.
func (g *GestionReservas) RegistrarReserva(r *Reserva) error {
	if err := r.Libro.ReducirStock(); err != nil {
		return err
	}
	g.lista = append(g.lista, r)
	return nil
}

func (g *GestionReservas) ListarReservas() []*Reserva {
	return g.lista
}
