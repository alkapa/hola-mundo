package randal

import (
	"fmt"
	"github.com/alkapa/hola-mundo/clima"
	"github.com/alkapa/hola-mundo/nacimiento"
	"time"
)

type (
	Randal struct {
		nacimiento.Nacimiento
	}
)

func NewRandal(nacimiento nacimiento.Nacimiento) Randal {
	return Randal{nacimiento}
}

func (*Randal) Saludo(como string, pregunta string) string {
	return "Hola " + como + " " + pregunta
}

func (*Randal) ComoEstaElClima(estacion clima.Estacion) string {

	if estacion == clima.Invierno {
		return "Esta malo el clima muy frio"
	}

	if estacion == clima.Verano {
		return "Esta bueno el clima para una fria"
	}

	return "Esta normal"
}

func (*Randal) EsTuCumpleanios(fecha string, nombre string) bool {
	return calculaCumpleanios(fecha, nombre)
}

func calculaCumpleanios(fecha string, nombre string) bool {
	miFechaDeNacimiento, err := time.Parse("02/01/2006", fecha)
	if err != nil {
		fmt.Println(err)
		return false
	}

	miNacimiento := nacimiento.NewNacimiento(miFechaDeNacimiento.Year(), miFechaDeNacimiento.Month(), miFechaDeNacimiento.Day())
	siCumpleAnios, diasRestantes := miNacimiento.Cumpleanios()
	if siCumpleAnios == false {
		fmt.Printf("Faltan %d dias para tu cumpleanios %s\n", diasRestantes, nombre)
	}

	return siCumpleAnios
}
