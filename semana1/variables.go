package main

import "fmt"

func main() {

	var nombre string = "Jorge" // ""
	var edad int
	var Apellido string = "Alvarez"

	edad = 26

	fmt.Println(nombre + Apellido)
	fmt.Println(edad)

	////parte 2

	otroNombre := "Pepe"
	otraEdad := 26

	fmt.Println(otroNombre + Apellido)
	fmt.Println(otraEdad)

	//parte 3

	var altura = 1.75
	fmt.Println(altura)

	//prueba
	var peso float32 = 75.8
	peso2 := 25.5

	var peso3 float32 = 76.9 //:= 28.33

	fmt.Println(" hola: ", peso, peso2, peso3)

	fmt.Println(" peso 1: ", peso, "peso 2:", peso2, "peso 3:", peso3)

}
