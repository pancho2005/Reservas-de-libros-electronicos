package main

import (
	"fmt"

	"biblioteca-go/libros"
	"biblioteca-go/reportes"
	"biblioteca-go/reservas"
	"biblioteca-go/usuarios"
)

func main() {

	usuario := usuarios.NewUsuario(1, "Francisco", "Piedra")
	libro := libros.NewLibro("Go Básico", "Autor X", 2)

	reserva := &reservas.Reserva{
		Usuario: usuario,
		Libro:   libro,
	}

	gestor := reservas.NewGestionReservas()

	if err := gestor.RegistrarReserva(reserva); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Reserva realizada para:", usuario.NombreCompleto())
	fmt.Println("Total de reservas:", reportes.TotalReservas(gestor.ListarReservas()))
}
