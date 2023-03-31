package main

import (
	"fmt"

	"github.com/harolpg17/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(145)
	fmt.Println(estado)
	fmt.Println(texto)
}
