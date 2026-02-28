package data

import (
	"biblioteca-go/models"
	"sync"
)

var Usuarios []models.Usuario
var Libros []models.Libro
var Reservas []models.Reserva

// Mutex para evitar condiciones de carrera
var Mutex sync.Mutex
