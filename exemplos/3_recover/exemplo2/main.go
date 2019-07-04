// tenta se recuperar dos erros

package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	executar()
}

func executar() {

	defer func() {
		if err := recover(); err != nil {
			log.Print("recuperando de um erro -> ", err)
			executar()
		}
	}()

	emitirNotaFiscal()
}

func emitirNotaFiscal() {
	// implementação da nota
	// emulando comunicação falha

	if fakeFalhou() {
		panic("nota fiscal não emitida: conexão falhou")
	}

	log.Print("Nota emitida com sucesso !!!")
}

func fakeFalhou() bool {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Int31n(10) != 5
}