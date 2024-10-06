package main

import "fmt"

func pi() float64 {
	return 3.1416
}

type CustomFunction func() float64

func main() {
	fmt.Println("Hello World")

	var mifuncion CustomFunction

	mifuncion = pi

	fmt.Println(mifuncion())

	fmt.Println("Bye World")
}
