package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() {
	scaner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error

	for {
		fmt.Println("Ingrese un número: ")
		if scaner.Scan() {
			numero, err = strconv.Atoi(scaner.Text())

			if err != nil {
				fmt.Println("Valor ingresado incorrecto, intenta de nuevo")
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(numero, "x", i, "=", numero*i)
	}

}
