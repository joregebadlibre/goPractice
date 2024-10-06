package main

import "fmt"

func mostrarVariable(objeto interface{}) {
	fmt.Println(objeto)

	fmt.Printf("eL VALOR ES DE TIPO: %T ", objeto)
}

func main() {

	mostrarVariable("hola")
	mostrarVariable(1)
	mostrarVariable(1.2)

}
