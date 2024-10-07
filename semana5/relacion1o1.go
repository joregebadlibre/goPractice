package main

import "fmt"

type user struct {
	FirstName string
	LastName  string
	Email     string
	Active    bool
}

type student struct {
	usuario user
	id      string
}

func main() {

	eduardo := user{
		FirstName: "Edgar",
		LastName:  "Alvarez",
		Email:     "Edgar@edgar",
		Active:    true,
	}

	ariel := user{
		FirstName: "Ariel",
		LastName:  "Alvarez",
		Email:     "Ariel@ariel",
		Active:    true,
	}

	eduardoStudent := student{
		usuario: eduardo,
		id:      "12345",
	}

	fmt.Println(eduardoStudent)
	fmt.Println(eduardoStudent.usuario.FirstName)
	fmt.Println(eduardoStudent.usuario.Email)
	fmt.Println(eduardoStudent.usuario.Active)
	fmt.Println(eduardoStudent.id)

	fmt.Println(ariel)
	fmt.Println(ariel.FirstName)
	fmt.Println(ariel.Email)
	fmt.Println(ariel.Active)

}
