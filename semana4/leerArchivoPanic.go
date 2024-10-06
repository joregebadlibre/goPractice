package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error al leer el archivo", err)
		}
	}()

	fmt.Println("Hello World")

	if archivo, err := os.Open("holas.txt"); err != nil {
		panic("Error al leer el archivo")
	} else {

		defer func() {
			err := archivo.Close()
			if err != nil {
				fmt.Println("Error al cerrar el archivo")
			}
		}()

		contenido := make([]byte, 254)
		longitud, _ := archivo.Read(contenido)
		contenidoArchivo := string(contenido[0:longitud])
		fmt.Println(contenidoArchivo)
	}
}
