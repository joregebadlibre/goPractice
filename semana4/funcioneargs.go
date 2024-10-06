package main

import "fmt"

type Operacion func(balances []int) int

func suma(numeros []int) int {
	resultado := 0
	for clave := range numeros {
		resultado += numeros[clave]
	}
	return resultado
}

func resta(numeros []int) int {
	resultado := 0
	for clave := range numeros {
		resultado -= numeros[clave]
	}
	return resultado
}

func multiplicar(numeros []int) int {
	resultado := 1
	for clave := range numeros {
		resultado *= numeros[clave]
	}
	return resultado
}

func division(numeros []int) int {
	resultado := 1
	for clave := range numeros {
		resultado /= numeros[clave]
	}
	return resultado
}

func operar(operacion Operacion, numeros []int) {
	fmt.Println("Antes de la operación: ")
	resultado := operacion(numeros)
	fmt.Println("El resultado es: ", resultado)
	fmt.Println("Despues de la operación")
}

func main() {

	operar(suma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	operar(resta, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	operar(multiplicar, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	operar(division, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
