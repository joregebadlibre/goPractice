package main

import "fmt"

func main() {

	monedas := [...]string{1: "sucre", 0: "peso", 2: "dolar", 3: "real"}

	fmt.Println(monedas)

	submonedas := monedas[0:2]
	fmt.Println(submonedas)

}
