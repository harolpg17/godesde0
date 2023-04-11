package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 9, 23, 5, 2, 4, 2, 23, 2, 32}

func MuestroSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:]   // Slice creado desde un vector desde la posición 3
	porcion2 := arreglo[:5]  // Slice creado desde un vector desde la posición 0 hasta la 5
	porcion3 := arreglo[6:8] // Slice creado desde un vector desde la posición 6 hasta la 8

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
