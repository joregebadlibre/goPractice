package main

import "fmt"

const (
	domingo int = iota + 1
	lunes
	martes
	miercoles
	jueves
	viernes
	sabado
)

func main() {

	for dia := domingo; dia <= sabado; dia++ {
		fmt.Println(dia)
	}

}
