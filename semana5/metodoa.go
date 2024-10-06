package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (self *User) setName(name string) {
	self.Name = name
}

func (self *User) getName() string {
	return self.Name
}

func main() {

	miuser := User{"Jorge", "Jorge@jorge"}
	fmt.Println(miuser.Name)
	miuser.setName("Jorge Alvarez")
	fmt.Println(miuser.getName())

}
