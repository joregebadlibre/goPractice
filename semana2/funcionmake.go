package main

import (
	"fmt"
)

func main() {

	numeros := make([]int, 3, 10)

	numeros[0] = 10
	numeros[1] = 100
	numeros[2] = 1000

	numeros = append(numeros, 10000, 100000)

	fmt.Println(numeros)
	fmt.Println(len(numeros))
	fmt.Println(cap(numeros))

}
