package main

import (
	"log"
	"onlineLibrary/internal/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("database connected successfully")
}
