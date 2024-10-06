package main

import "fmt"

type MIOperacion func(balances []int) int

func crearOperacion(tipoOperacion string) MIOperacion {

	if tipoOperacion == "suma" {
		return func(balances []int) int {
			return suma(balances)
		}
	} else if tipoOperacion == "resta" {
		return func(balances []int) int {
			return resta(balances)
		}
	}
	return nil //no se ejecuta nunca
}

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

func main() {

	operacion := crearOperacion("suma")
	resultado := operacion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("El resultado es: ", resultado)

	operacion = crearOperacion("resta")
	resultado = operacion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("El resultado es: ", resultado)
}
