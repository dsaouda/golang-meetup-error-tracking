// fazer alguma coisa quando um erro ocorrer

package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			// enviar um e-mail com problema
			// gerar log
		}
	}()

	emitirNotaFiscal()
}

func emitirNotaFiscal() {
	// implementação da nota
	// emulando comunicação falha
	panic("nota fiscal não emitida: conexão falhou")
}

