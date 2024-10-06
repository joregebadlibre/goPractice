package main

import "fmt"

func main() {

	miFuncion := func(username string) string {
		message := "Hola: " + username
		return message
	}
	resultado := miFuncion("Jorge")
	fmt.Println("El resultado es: ", resultado)
	fmt.Println(miFuncion("JorgeR"))
}
