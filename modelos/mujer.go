package modelos

type Mujer struct {
	Hombre
}

func (mujer *Mujer) Sexo() string {
	return "Mujer"
}
