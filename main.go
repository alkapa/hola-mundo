package main

import (
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
}
