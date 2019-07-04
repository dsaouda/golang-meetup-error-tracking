package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal("erro ao conectar: ", err.Error())
	}
	defer db.Close()

	// realizar consultas ...
}
