package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("chamada 1")
	defer fmt.Println("chamada 2")
	defer fmt.Println("chamada 3")
	defer fmt.Println("chamada 4")
	defer fmt.Println("chamada 5")

	fmt.Println("chamada 6")
}
