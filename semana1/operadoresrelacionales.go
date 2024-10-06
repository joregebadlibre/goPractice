package main

import "fmt"

func main() {

	/*
	   <
	   >
	   <=
	   >=
	   !=
	   ==
	*/

	var edad = 10
	var resultado = edad >= 18
	fmt.Println("es mayor de edad", resultado)

	var nombre = "Jorge"
	var otroNombre = "Pepe"
	var resultado2 = nombre != otroNombre

	fmt.Println("es diferente", resultado2)

}
