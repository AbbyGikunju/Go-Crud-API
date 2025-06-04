package main

import (
	"fmt"
	"go-crud-api/db"

	
)

func main() {
	db.INITPostgresDB()
	fmt.Println("Connected to the database")
}