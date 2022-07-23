package main

import (
	"fmt"
	"github.com/alkapa/hola-mundo/clima"
	"github.com/alkapa/hola-mundo/randal"
)

func main() {
	println(
		randal.Saludo("primo", "como estas?"),
	)
	println(
		randal.ComoEstaElClima(clima.ObtenerEstacionActual()),
	)

	if randal.EsMiCumpleanios() {
		fmt.Println("feliz cumpleanios para mi")
	}

	/*
		TODO: crear una tipo de dato nuevo llamado Amigo
		que va a tener 2 campos
			- Nombre de tipo string
			- Fecha de Nacimiento de tipo string

		Buscar como crear un type struct para solucionar  esta tarea
	*/

	if randal.EsTuCumpleanios("27/06/1993") {
		fmt.Printf("feliz cumple %s", "camilo")
	}
}
