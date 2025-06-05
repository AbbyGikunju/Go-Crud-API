package main

import (
	"go-crud-api/db"
	"go-crud-api/router"
)

func main() {
	db.INITPostgresDB()
	router.InitRouter().Run()
}