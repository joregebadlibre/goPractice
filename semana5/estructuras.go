package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

func main() {

	//user := User{"Jorge", "Jorge@jorge"}
	user := User{Name: "Jorge", Email: "Jorge@jorge"}

	fmt.Println(user.Name)
	fmt.Println(user.Email)

}
