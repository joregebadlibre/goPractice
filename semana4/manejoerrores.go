package main

import (
	"errors"
	"fmt"
)

func midivision(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return a / b, nil
}

func main() {

	if resultado, err := midivision(10, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El resultado es: ", resultado)
	}
}
