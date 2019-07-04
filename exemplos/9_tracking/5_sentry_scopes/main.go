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

	// global
	//sentry.ConfigureScope(func(scope *sentry.Scope) {
	//
	//	// user
	//	scope.SetUser(sentry.User{Email: "john.doe@example.com"})
	//
	//	// tag
	//	scope.SetTag("page_locale", "pt-br")
	//
	//	// extra
	//	scope.SetExtra("character_name", "Mighty Fighter")
	//})

	_, err := os.Open("filename2.ext")
	if err != nil {
		// local
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetUser(sentry.User{Email: "john.doe@example.com"})
			scope.SetTag("page_locale", "pt-br")
			scope.SetExtra("character_name", "Mighty Fighter")
			sentry.CaptureException(err)
		})
	}

	sentry.Flush(time.Second*10)
}
