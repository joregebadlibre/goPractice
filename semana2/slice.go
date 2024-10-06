package main

import "fmt"

func main() {
	//slices

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numeros = append(numeros, 888, 999)
	numeros = append(numeros, 777)

	nuevoSlice := numeros[0:5]
	tecerSlice := numeros[5:10]

	numeros = append(numeros, nuevoSlice[0:1]...)
	numeros = append(numeros, tecerSlice...)

	numeros[0] = 2222
	numeros[9] = 3333

	fmt.Println(numeros)
	fmt.Println(nuevoSlice)
	fmt.Println(tecerSlice)

}
