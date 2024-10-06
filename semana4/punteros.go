package main

import "fmt"

func mimodificar(parametro *string) {
	*parametro = "Cambia de valor"
}

func main() {

	//Punteros
	//Punteros son variables que guardan una dirección de memoria
	//Punteros son muy útiles para pasar valores por referencia

	var variable string = "Jorge"

	fmt.Println("Valor actual", variable)
	mimodificar(&variable)
	fmt.Println("Valor modificado", variable)
}
