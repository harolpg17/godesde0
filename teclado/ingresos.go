package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scaner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1: ")
	if scaner.Scan() {
		numero1, err = strconv.Atoi(scaner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto: " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2: ")
	if scaner.Scan() {
		numero2, err = strconv.Atoi(scaner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto: " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scaner.Scan() {
		leyenda = scaner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
