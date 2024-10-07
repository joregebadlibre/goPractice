package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var listusuarios map[int]User
var id int = 0

func crearusuario() {
	clearConsole()
	id++

	nuevoUsuario := formulario("Ingresar")
	nuevoUsuario.id = id
	listusuarios[id] = nuevoUsuario
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>Usuario creado \n")
}

func listarusuario() {
	clearConsole()
	fmt.Println("Queriendo listar usuario")
	for clave, valor := range listusuarios {
		fmt.Println(clave, "-", valor.username)
	}
}

func formulario(accion string) User {

	fmt.Print(accion, " el nuevo nombre: ")
	username := readLine()
	fmt.Print(accion, " el nuevo email: ")
	email := readLine()
	fmt.Print(accion, " la nueva edad: ")
	edad := readLineInt()
	return User{
		username: username,
		email:    email,
		age:      edad,
	}
}

func actualizarusuario() {
	clearConsole()
	fmt.Print("Ingresa el Id del usuario que quieres actualizar: ")
	id := readLineInt()
	if buscarusuario(id) {
		listusuarios[id] = formulario("actualiza")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>Usuario actualizado \n")
	}
}

func buscarusuario(id int) bool {
	if _, ok := listusuarios[id]; ok {
		fmt.Println("Usuario encontrado")
		return true
	} else {
		fmt.Println("Usuario no encontrado")
		return false
	}
}

func eliminarusuario() {
	clearConsole()
	fmt.Print("Ingresa el Id del usuario que quieres eliminar: ")
	id := readLineInt()
	if buscarusuario(id) {
		delete(listusuarios, id)
		fmt.Println("Usuario eliminado")
	}
}

func imprimirMenu() {
	fmt.Println("Bienvenido al sistema de usuarios")
	fmt.Println("A) crear")
	fmt.Println("B) listar")
	fmt.Println("C) actualizar")
	fmt.Println("D) eliminar")
	fmt.Println("Q) Salir")
	fmt.Println("Ingresa una opción: ")
}

func readLine() string {
	reader = bufio.NewReader(os.Stdin)
	if option, err := reader.ReadString('\n'); err != nil {
		panic("Error al leer la entrada: ")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func readLineInt() int {
	reader = bufio.NewReader(os.Stdin)
	if option, err := reader.ReadString('\n'); err != nil {
		panic("Error al leer la entrada: ")
	} else {
		if optionInt, err := strconv.Atoi(strings.TrimSuffix(option, "\n")); err != nil {
			panic("Error al convertir la entrada: ")
		} else {
			return optionInt
		}
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error al leer la entrada: ", err)
		}
	}()

	listusuarios = make(map[int]User)
	for {
		imprimirMenu()
		option := readLine()
		switch strings.ToLower(option) {
		case "a", "crear":
			crearusuario()
		case "b", "listar":
			listarusuario()
		case "c", "actualizar":
			actualizarusuario()
		case "d", "eliminar":
			eliminarusuario()
		case "q", "salir":
			fmt.Println("Bye World")
			return
		default:
			fmt.Println("Opción incorrecta: ", option)
		}
	}
}
