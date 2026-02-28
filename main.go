package main

import (
	"biblioteca-go/handlers"
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// Servir página web
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	// Rutas API
	mux.HandleFunc("/usuarios", handlers.CrearUsuario)
	mux.HandleFunc("/usuarios/listar", handlers.ListarUsuarios)

	mux.HandleFunc("/libros", handlers.CrearLibro)
	mux.HandleFunc("/libros/listar", handlers.ListarLibros)

	mux.HandleFunc("/reservas", handlers.CrearReserva)
	mux.HandleFunc("/reservas/listar", handlers.ListarReservas)

	mux.HandleFunc("/reportes/total", handlers.TotalReservas)
	mux.HandleFunc("/reportes/libros", handlers.LibrosDisponibles)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
