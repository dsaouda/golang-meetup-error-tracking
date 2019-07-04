package main

import (
	"github.com/getsentry/sentry-go"
	"time"
)

func main() {

	_ = sentry.Init(sentry.ClientOptions{
		SampleRate: 1.0,
		// sentry informa dsn por env. Crie a env SENTRY_DSN=<seu_dsn>
		//Dsn:              "<dsn>",
		AttachStacktrace: true,
		Debug: true,
		Environment: "production",
	})

	sentry.CaptureMessage("mensagem de producao")

	_ = sentry.Init(sentry.ClientOptions{
		SampleRate: 1.0,
		// sentry informa dsn por env. Crie a env SENTRY_DSN=<seu_dsn>
		//Dsn:              "<dsn>",
		AttachStacktrace: true,
		Debug: true,
		Environment: "dev",
	})

	sentry.CaptureMessage("mensagem dev")

	sentry.Flush(time.Second*10)
}
