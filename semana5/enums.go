package main

import "fmt"

type UserType int

const (
	Teacher UserType = iota
	Student
	Admin
)

type User struct {
	Usename string
	Type    UserType
}

func main() {

	eduardo := User{
		Usename: "Edgar",
		Type:    Teacher,
	}
	ariel := User{
		Usename: "Ariel",
		Type:    Teacher,
	}

	fmt.Println(eduardo)
	fmt.Println(ariel)

	if eduardo.Type == Teacher {
		fmt.Println("Es un profesor")
	}

}
