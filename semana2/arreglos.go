package main

import "fmt"

func main() {

	var numeros [5]int = [5]int{100, 2, 3, 4, 5}

	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println(numeros)

	numeros2 := [5]int{1000, 2000, 3000, 4000, 5000}
	fmt.Println(numeros2)

	numerossindefinir := [...]int{101, 201, 301, 401, 501}
	fmt.Println(numerossindefinir)

}
