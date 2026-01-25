package main

import "fmt"

type Libro struct {
	ID           int
	Titulo       string
	Autor        string
	EstaPrestado bool
}

type Persona struct {
	ID       int
	Nombre   int
	Apellido string
}

func main() {
	// 1. Declaramos las variables donde guardaremos los datos
	var nombre string
	var edad int

	// 2. Pedimos el nombre (Usamos Print sin 'ln' para que el cursor se quede al lado)
	fmt.Print("¿Cómo te llamas?: ")

	// 3. ESCUCHAMOS al usuario
	// Importante: &nombre le dice a Go "guarda el dato en esta dirección de memoria"
	fmt.Scanln(&nombre)

	// 4. Pedimos la edad
	fmt.Print("¿Cuántos años tienes?: ")
	fmt.Scanln(&edad)

	// 5. Mostramos todo junto
	fmt.Println("------------------------------")
	fmt.Println("¡Genial! Te llamas", nombre, "y tienes", edad, "años.")

	// Usamos un condicional simple de lo que aprendimos antes
	if edad >= 18 {
		fmt.Println("Eres mayor de edad, ¡a programar se ha dicho!")
	} else {
		fmt.Println("Eres joven, ¡tienes mucho futuro por delante!")
	}
}
