package main

import (
	"fmt"
	"reflect"
)

func main() {

	var curso string = "Curso go"
	fmt.Println(curso)

	var edad = 15
	fmt.Println(edad)

	apellido := "Alvarez"
	fmt.Println(apellido)

	fmt.Println(len(curso), len(apellido))

	primercaracter := apellido[0]
	ultimocaracter := apellido[len(apellido)-1]

	fmt.Println(primercaracter)
	fmt.Println(string(primercaracter))
	fmt.Println(reflect.TypeOf(primercaracter))

	fmt.Printf("%c\n", primercaracter)

	fmt.Println(ultimocaracter)
	fmt.Println(string(ultimocaracter))
	fmt.Println(reflect.TypeOf(ultimocaracter))

	fmt.Printf("%c\n", ultimocaracter)

}
