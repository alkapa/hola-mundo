package main

import (
	"fmt"
	"github.com/alkapa/hola-mundo/camilo"
	"github.com/alkapa/hola-mundo/clima"
	"github.com/alkapa/hola-mundo/nacimiento"
	"github.com/alkapa/hola-mundo/randal"
)

func main() {
	yo := randal.NewRandal(nacimiento.NewNacimiento(1997, 7, 26))
	cami := camilo.NewCamilo(nacimiento.NewNacimiento(1997, 7, 26))

	println(
		yo.Saludo("primo", "como estas?"),
	)
	println(
		yo.ComoEstaElClima(clima.ObtenerEstacionActual()),
	)

	if ok, _ := yo.Cumpleanios(); ok {
		fmt.Println("feliz cumpleanios para mi")
	}

	if ok, _ := cami.Cumpleanios(); ok {
		fmt.Println("feliz cumpleanios para cami")
	}

	// TODO 4: convertir Amigo a Persona
	type Amigo struct {
		Nombre            string
		FechaDeNacimiento string
	}

	mAN := "Camilo"
	mAF := "27/06/1993"

	if yo.EsTuCumpleanios(mAF, mAN) {
		return
	}
}
