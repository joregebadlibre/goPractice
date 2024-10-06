package main

import "fmt"

func main() {

	/*
	   &&
	   ||
	   !
	*/

	resultado := 20 == 20 && 30 == 30
	fmt.Println("1: es verdadero", resultado)

	resultado2 := 20 > 20 || 30 == 15
	fmt.Println("2: es falso", resultado2)

	resultado3 := !(20 == 20)
	fmt.Println("4: es falso", resultado3)

	resultado4 := !(20 == 20 && 30 == 30) || (1 == 1)
	fmt.Println("4: es verdadero", resultado4)

}
