package main

import "fmt"

func modificar(variable string) {

	fmt.Println("Valor de mi variable", variable)
	variable = "Pepe"
	fmt.Println("Valor de mi variable", variable)

	fmt.Printf("%p\n", &variable)
}

func main() {

	mivariable := "Jorge"
	modificar(mivariable)
	fmt.Println("Valor de mi variable: ", mivariable)

	fmt.Printf("%p\n", &mivariable)
}
