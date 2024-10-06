package main

import "fmt"

func func1() {
	fmt.Println("Funcion1: ")
}

func func2() {
	fmt.Println("Funcion2: ")
}

func main() {
	fmt.Println("Hello World")

	defer func1()
	func2()

}
