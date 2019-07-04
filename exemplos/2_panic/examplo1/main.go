package main

import (
	"os"
	"strings"
)

func main() {

	db := os.Getenv("DB_HOST")

	if strings.TrimSpace(db) == "" {
		panic("não é possível encontrar a variável de ambiente DB_HOST")
		//log.Panicf("não é possível encontrar a variável de ambiente: %v", "DB_HOST")
	}

	// implementacao de conexão com banco de dados


}

