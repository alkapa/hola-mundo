package clima

import (
	"time"
)

type Estacion int

const (
	Invierno Estacion = iota
	Verano
	Otonio
	Primavera
	NONE
)

func ObtenerEstacionActual() Estacion {
	list := map[time.Month]Estacion{}

	for i := 1; i < 13; i++ {

		if i <= 3 {
			list[time.Month(i)] = Verano
			continue
		}

		if i > 3 && i <= 6 {
			list[time.Month(i)] = Otonio
			continue
		}

		if i > 6 && i <= 9 {
			list[time.Month(i)] = Invierno
			continue
		}

		if i > 9 && i <= 12 {
			list[time.Month(i)] = Primavera
			continue
		}
	}

	month := time.Now().Month()

	if v, ok := list[month]; ok {
		return v
	}

	return NONE
}
