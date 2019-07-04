package main

import (
	"fmt"
	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "O arquivo não foi localizado")
	}

	return f, nil
}
