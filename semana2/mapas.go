package main

import (
	"fmt"
)

func main() {

	//mapas

	dias := make(map[int]string)
	dias[0] = "domingo"
	dias[1] = "lunes"
	dias[2] = "martes"
	dias[3] = "miercoles"
	dias[4] = "jueves"
	dias[5] = "viernes"
	dias[6] = "sabado"

	dias[4] = "JUEVES"

	delete(dias, 4)

	fmt.Println(dias)

	usuarios := make(map[string][]int)
	usuarios["Jorge"] = []int{1, 2, 3}
	usuarios["Pepe"] = []int{4, 5, 6}

	fmt.Println(usuarios)
}
