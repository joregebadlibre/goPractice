package main

import "fmt"

func factorial(numero int) int {
	if numero == 1 {
		return 1
	}
	return numero * factorial(numero-1)
}

func main() {
	fmt.Println("Hello World")

	resultado := factorial(5)

	fmt.Println("El factorial es: ", resultado)

}
