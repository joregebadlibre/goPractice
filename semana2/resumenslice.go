package main

import "fmt"

func main() {

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numeros[0] = 111
	numeros[9] = 222

	inicio := numeros[0:3]
	final := numeros[6:10]

	fmt.Println(numeros)
	fmt.Println(inicio)
	fmt.Println(final)

}
