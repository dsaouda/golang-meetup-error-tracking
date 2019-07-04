package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"
	"os"
	"time"
)

func main() {
	_ = sentry.Init(sentry.ClientOptions{
		SampleRate: 1,
		// sentry informa dsn por env. Crie a env SENTRY_DSN=<seu_dsn>
		//Dsn:              "<dsn>",
		AttachStacktrace: true,
	})

	sentry.CaptureException(chamada1())
	sentry.Flush(time.Second * 10)
}

func chamada1() error {
	return errors.WrapPrefix(chamada2(), "A chamada 2 retornou erro", 0)
}

func chamada2() error {
	return chamada3()
}

func chamada3() error {
	return chamada4()
}

func chamada4() error {

	_, err := os.Open("arquivo_nao_existe")

	return errors.WrapPrefix(err, "(Erro ao abrir o arquivo)", 0)
}
