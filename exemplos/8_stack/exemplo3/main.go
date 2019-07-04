package main

import (
	"fmt"
	"github.com/pkg/errors" //com stack
	//"errors" //sem stack

)

func main() {
	fmt.Printf("%+v", chamada1())
}

func chamada1() error {
	return chamada2()
}

func chamada2() error {
	return chamada3()
}

func chamada3() error {
	return chamada4()
}

func chamada4() error {
	return errors.New("Ocorreu um erro na chamada 4")
}
