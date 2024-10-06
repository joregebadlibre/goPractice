package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World")
	file, err := os.Open("hola.txt")

	if err != nil {
		fmt.Println("Error al leer el archivo")
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Error al cerrar el archivo")
		}
	}()

	contenido := make([]byte, 254)
	longitud, _ := file.Read(contenido)
	contenidoArchivo := string(contenido[0:longitud])
	fmt.Println(contenidoArchivo)
}
