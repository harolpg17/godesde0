package main

import (
	"fmt"
	"runtime"

	"github.com/harolpg17/godesde0/arreglos_slices"
	"github.com/harolpg17/godesde0/ejer_interfaces"
	"github.com/harolpg17/godesde0/ejercicios"
	"github.com/harolpg17/godesde0/files"
	"github.com/harolpg17/godesde0/funciones"
	"github.com/harolpg17/godesde0/iteraciones"
	"github.com/harolpg17/godesde0/mapas"
	"github.com/harolpg17/godesde0/modelos"
	"github.com/harolpg17/godesde0/users"
	"github.com/harolpg17/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	fmt.Println("-----------------------------")

	variables.RestoVariables()
	fmt.Println("-----------------------------")

	estado, texto := variables.ConviertoaTexto(145)
	fmt.Println(estado)
	fmt.Println(texto)
	fmt.Println("-----------------------------")

	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Esto es Linux, es ", os)
	} else {
		fmt.Println("Esto no es Linux, es ", os)
	}
	fmt.Println("-----------------------------")

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "windows":
		fmt.Println("Esto es Windows")
	default:
		fmt.Printf("%s 	\n", os)
	}

	fmt.Println("-----------------------------")

	numero, mensaje := ejercicios.ConvertirTextoaEntero("60")
	fmt.Println(numero)
	fmt.Println(mensaje)
	fmt.Println("-----------------------------")

	// teclado.IngresoNumeros()
	fmt.Println("-----------------------------")

	iteraciones.Iterar()
	fmt.Println("-----------------------------")

	// fmt.Println(ejercicios.TablaDeMultiplicar())

	fmt.Println("-----------------------------")

	// files.GrabaTabla()

	fmt.Println("-----------------------------")

	files.SumaTabla()
	fmt.Println("-----------------------------")

	files.LeoArchivo()

	fmt.Println("-----------------------------")

	funciones.Calculos()

	fmt.Println("-----------------------------")

	funciones.LamarClosures()

	fmt.Println("-----------------------------")

	funciones.Exponencia(2)

	fmt.Println("-----------------------------")

	arreglos_slices.MuestroArreglos()

	arreglos_slices.MuestroSlice()

	arreglos_slices.Capacidad()

	fmt.Println("-----------------------------")
	mapas.MostrarMapas()

	fmt.Println("-----------------------------")

	users.AltaUsuario()

	fmt.Println("-----------------------------")

	Pedro := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Maria)
}
