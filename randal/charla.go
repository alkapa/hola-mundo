package randal

import (
	"fmt"
	"github.com/alkapa/hola-mundo/clima"
	"time"
)

func Saludo(como string, pregunta string) string {
	return "Hola " + como + " " + pregunta
}

func ComoEstaElClima(estacion clima.Estacion) string {

	if estacion == clima.Invierno {
		return "Esta malo el clima muy frio"
	}

	if estacion == clima.Verano {
		return "Esta bueno el clima para una fria"
	}

	return "Esta normal"
}

func EsMiCumpleanios() bool {
	const cumpleanios = "02/11/1997"
	// TODO: agregar constante con tu nombre y pasarla a la nueva firna de la funcion
	return calculaCumpleanios(cumpleanios)
}

/*
	TODO: agregar el nombre de la persona a la firma de la funcion
*/
func EsTuCumpleanios(fecha string) bool {
	return calculaCumpleanios(fecha)
}

/*
	TODO: agregar el nombre de la persona a la firma de la funcion
*/
func calculaCumpleanios(fecha string) bool {
	micumpleanios, err := time.Parse("02/01/2006", fecha)
	if err != nil {
		fmt.Println(err)
		return false
	}

	diaCumple := micumpleanios.Day()
	mesCumple := micumpleanios.Month()

	fechaActual := time.Now()
	fechaActual = time.Date(fechaActual.Year(), fechaActual.Month(), fechaActual.Day(), 0, 0, 0, 0, time.Local)

	anioActual := fechaActual.Year()

	cumpleActual := time.Date(anioActual, mesCumple, diaCumple, 0, 0, 0, 0, time.Local)

	if fechaActual.After(cumpleActual) {
		cumpleActual = cumpleActual.AddDate(1, 0, 0)
	}

	diferencia := cumpleActual.Sub(fechaActual)
	diasRestantes := diferencia.Hours() / 24

	if diferencia != 0 {
		// TODO: agregar el nombre de la persona y colocarlo dentro del mensaje  para saber a quien le preguntamos
		fmt.Printf("te faltan %f dias para mi cumpleanios :nombre:", diasRestantes)
		return false
	}

	return true
}
