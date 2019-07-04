package main

import (
	"fmt"
	"os"
)

func main() {

	a, err := abrirArquivo("arquivo_nao_existe")

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(0)
	}

	defer a.Close()
}

func abrirArquivo(filename string) (*os.File, error) {

	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return f, nil
}
