package camilo

import "github.com/alkapa/hola-mundo/nacimiento"

type (
	Camilo struct {
		nacimiento.Nacimiento
	}
)

func NewCamilo(nacimiento nacimiento.Nacimiento) Camilo {
	return Camilo{nacimiento}
}
