package main

import "fmt"

func promedio(calificaciones ...int) int {
	var total float64 = 0
	for _, calificacion := range calificaciones {
		total += float64(calificacion)
		fmt.Println(calificacion)
	}
	return int(total / float64(len(calificaciones)))
}

func main() {
	fmt.Println("Hello World")

	resultado := promedio(6, 7, 8, 9, 10)

	fmt.Println("El promedio es: ", resultado)

}
