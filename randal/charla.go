package randal

import "github.com/alkapa/hola-mundo/clima"

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

func MiCumpleanios() bool {
	/*
		1- definir una constante con la fecha de tu cumpleanios
		2- usar esa constante e imprimir cuantos dias te faltan para tu proximo cumple
		3- usar el paquete time para la fecha y sus calculos
		4- usar una cadena de caracters para la fecha y convertirla en time.Time
		5- retornar verdadero (true) si el dia de hoy es tu cumpleanios sino retornar falso
		   (false) eh imprimir cuantos dias faltan
	*/
	panic("borrame cuando termines la implementacion")
}
