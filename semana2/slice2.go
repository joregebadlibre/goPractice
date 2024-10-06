package main

import "fmt"

func main() {
	//slices
	// Un puntero
	//Una longitud
	//Una capacidad

	meses := make([]string, 0, 12)
	meses = append(meses, "enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre")
	meses = []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}
	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Println(meses)

	fmt.Println("longitud:", longitud)
	fmt.Println("capacidad:", capacidad)

	meses = append(meses, "Octobre")
	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Println(meses)
	fmt.Println("longitud:", longitud)
	fmt.Println("capacidad:", capacidad)

}
