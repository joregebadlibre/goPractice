package main

import (
	"log"
	"pruebapostgres/storage"
)

func main() {
	log.Println("Start Project")

	DB := storage.ConnectDB()
	defer DB.Close()

	log.Println(DB.RowsAffected)
}
