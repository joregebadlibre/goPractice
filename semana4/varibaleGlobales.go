package main

import "fmt"

var username = ""

func func1() {
	fmt.Println("Funcion1: ", username)
}

func func2() {
	fmt.Println("Funcion2: ", username)
	username = "Funcion 2"
}

func main() {

	func1()
	func2()
	username = "Funcion Principal"
	func2()
	func1()
}
