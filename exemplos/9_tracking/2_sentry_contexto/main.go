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
	})

	// extra ...
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetExtra("character_name", "Mighty Fighter")
		sentry.CaptureMessage("enviando extra")
	})

	// user ...
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{Email: "john.doe@example.com"})
		sentry.CaptureMessage("enviando user")
	})

	// tag ...
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetTag("page_locale", "pt-br")
		sentry.CaptureMessage("enviando tag")
	})

	sentry.Flush(time.Second*10)
}
