package main

import "fmt"

const (
	nombre   string = "Jorge"   //constante
	edad     int    = 26        //constante
	apellido string = "Alvarez" //constante
)

const curso string = "Curso de Go"

func main() {

	fmt.Println(curso)
	fmt.Println(nombre, edad, apellido)

}
