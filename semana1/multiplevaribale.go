package main

import "fmt"

func main() {

	newFunction()

}

func newFunction() {
	var nombre string = "Jorge"
	var apellido string = "Alvarez"
	var pais string = "Ecuador"

	var nombre2, apellido2, pais2 string = "Jorge", "Alvarez", "Ecuador"

	fmt.Println(nombre2, apellido2, pais2)

	fmt.Println(nombre, apellido, pais)

	nombre3, apellido3, pais3 := "Jorge", "Alvarez", "Ecuador"

	fmt.Println(nombre3, apellido3, pais3)
}
