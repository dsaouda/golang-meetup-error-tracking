package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("nao_existe")

	if err != nil {
		fmt.Println(err)
		//fmt.Println(fmt.Sprintf("ocorreu um erro ao abrir arquivo (%v)", err))
		os.Exit(0)
	}

	defer f.Close()

}
