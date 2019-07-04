package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Open("nao_existe")

	if err != nil {
		log.Fatalf("ocorreu um erro ao abrir arquivo (%v)", err)
	}

	defer f.Close()

}
