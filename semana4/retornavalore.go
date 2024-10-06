package main

import "fmt"

func suma(numero1 int, numero2 int) (string, int) {
	return "La suma es: ", numero1 + numero2
}

func main() {

	fmt.Println(suma(10, 20))

	mensaje, resultado := suma(10, 20)
	fmt.Println(mensaje, resultado)
}
