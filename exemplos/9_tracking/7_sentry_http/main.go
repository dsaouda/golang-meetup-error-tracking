package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)
func main() {
	_ = sentry.Init(sentry.ClientOptions{
		// sentry informa dsn por env. Crie a env SENTRY_DSN=<seu_dsn>
		//Dsn:              "<dsn>",
		Debug:            true,
		AttachStacktrace: true,
	})

	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	// forçar divisão por 0
	http.HandleFunc("/divisao-por-zero", sentryHandler.HandleFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			a := 10
			b, _ := strconv.Atoi("0")
			_ = a/b
		},
	))

	// panic explícito
	http.HandleFunc("/forcado", sentryHandler.HandleFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			panic("erro forcado")
		},
	))

	fmt.Println("Listening and serving HTTP on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}