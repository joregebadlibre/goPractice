package main

import (
	"fmt"

	codigof "github.com/joregebadlibre/goPractice/semana5/codigof"
)

func main() {

	curso := codigof.Curso{"Curso de Go"}
	fmt.Println(curso.ObtenerMetodo())
}
