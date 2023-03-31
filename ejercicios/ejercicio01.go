package ejercicios

import (
	"strconv"
)

func ConvertirTextoaEntero(numeroEnTexto string) (int, string) {
	numero, error := strconv.Atoi(numeroEnTexto)
	if error != nil {
		return 0, "Hubo un error " + error.Error()
	}

	var mensaje string
	if numero > 100 {
		mensaje = "Es mayor a 100"
	} else {
		mensaje = "Es menor a 100"
	}

	return numero, mensaje
}
