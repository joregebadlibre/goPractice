package main

import "fmt"

type Animal interface {
	Comer()
	Dormir()
	Cazar()
}

type Perro struct {
	nombre string
}

func (p Perro) Comer() {
	fmt.Println("El perro come")
}

func (p Perro) Dormir() {

	fmt.Println("El perro duerme")
}

func (p Perro) Cazar() {
	fmt.Println("El perro caza")
}

func ejecutarAccion(animal Animal) {
	animal.Comer()
	animal.Dormir()
	animal.Cazar()
}

func main() {

	loki := Perro{"Loki"}
	ejecutarAccion(loki)

}
