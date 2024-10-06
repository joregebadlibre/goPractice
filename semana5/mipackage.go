package main

import (
	"fmt"

	"./codigof"
)

func main() {

	var curso Curso = codigof.Curso{titulo: "Curso de Go"}

	fmt.Println(curso)

}
