package main

import "fmt"

type Animal interface {
	Dormir()
}
type Mascota interface {
	Comer()
}

type Gato struct {
	nombre string
}

func (g Gato) Dormir() {
	fmt.Println("El gato duerme")
}

func (g Gato) Comer() {
	fmt.Println("El gato come")
}

func funcionParaunAnimal(animal Animal) {
	fmt.Println("Funcion para un animal")
}

func funcionParaunaMascota(mascota Mascota) {
	fmt.Println("Funcion para una mascota")
}

func main() {

	gato := Gato{nombre: "mi gato"}
	gato.Dormir()
	gato.Comer()

	funcionParaunAnimal(gato)
	funcionParaunaMascota(gato)
}
