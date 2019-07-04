package main

import (
	"github.com/getsentry/sentry-go"
	"os"
	"time"
)

func main() {

	_ = sentry.Init(sentry.ClientOptions{
		SampleRate: 1.0,
		// sentry informa dsn por env. Crie a env SENTRY_DSN=<seu_dsn>
		//Dsn:              "<dsn>",
		AttachStacktrace: true,
		Debug: true,
	})

	_, err := os.Open("filename.ext")
	if err != nil {
		sentry.CaptureException(err)
	}

	sentry.Flush(time.Second*10)
}
