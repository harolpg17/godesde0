package main

import (
	"fmt"
	"runtime"

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

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "windows":
		fmt.Println("Esto es Windows")
	default:
		fmt.Printf("%s 	\n", os)
	}
}
