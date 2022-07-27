package nacimiento

import (
	"time"
)

type (
	Nacimiento struct {
		Anio int
		Mes  time.Month
		Dia  int
	}
)

func NewNacimiento(yyyy int, mm time.Month, dd int) Nacimiento {
	return Nacimiento{
		Anio: yyyy,
		Mes:  mm,
		Dia:  dd,
	}
}

func (n *Nacimiento) Cumpleanios() (siCumpleAnios bool, diasRestantes int) {
	fechaActual := time.Now()
	fechaActual = time.Date(fechaActual.Year(), fechaActual.Month(), fechaActual.Day(), 0, 0, 0, 0, time.Local)

	anioActual := fechaActual.Year()

	cumpleActual := time.Date(anioActual, n.Mes, n.Dia, 0, 0, 0, 0, time.Local)

	if fechaActual.After(cumpleActual) {
		cumpleActual = cumpleActual.AddDate(1, 0, 0)
	}

	diferencia := cumpleActual.Sub(fechaActual)

	if diferencia != 0 {
		return false, int(diferencia.Hours() / 24)
	}

	return true, 0
}
