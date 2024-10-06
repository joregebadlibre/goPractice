package main

import "fmt"

func main() { //Bloque 1

	var curso = "Curso de Go"
	var version = 1.1
	{ //Bloque 2

		fmt.Println(curso)
		nuevoCurso := "Curso de Python"
		fmt.Println(nuevoCurso)
		{ //Bloque 3

			fmt.Println(curso)
			tercerCurso := "Curso de JavaScript"
			fmt.Println(tercerCurso)
			var version = 1.2
			fmt.Println("Version: ", version)
		}
	}
	fmt.Println("Curso: ", curso)
	fmt.Println("Version: ", version)
}
