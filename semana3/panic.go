package main

import (
	"fmt"
)

func main() {
	//panic
	var dividendo, divisor int = 10, 2

	fmt.Print("Ingresa el dividendo: ")
	fmt.Scanf("%d", &dividendo)

	if divisor == 0 {
		panic("No se puede dividir por cero")
	}

	resultado := dividendo / divisor
	fmt.Println("El resultado es: ", resultado)
}
