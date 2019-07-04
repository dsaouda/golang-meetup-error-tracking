package main

import (
	"os"
	"palestra/10_abstrair_erros/meuerro"
	"time"
)

func main() {

	if err := call1(); err != nil {
		meuerro.Notify(err)
	}

	time.Sleep(time.Second*10)

}

func call1() error {
	return meuerro.ArgumentWithError(call2())
}

func call2() error {
	return call3()
}

func call3() error {
	return exemploErro()
}

func exemploErro() error {
	_, err := os.Open("nao_existe")
	return meuerro.Wrap(err, "exemplo com erro")
}