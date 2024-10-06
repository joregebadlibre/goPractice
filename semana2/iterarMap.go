package main

import "fmt"

func main() {
	//mapas

	usuarios := map[int]string{}

	usuarios[1] = "Jorge"
	usuarios[3] = "Pepe"
	usuarios[2] = "Maria"

	for clave, valor := range usuarios {

		fmt.Println(clave, valor)
	}
}
