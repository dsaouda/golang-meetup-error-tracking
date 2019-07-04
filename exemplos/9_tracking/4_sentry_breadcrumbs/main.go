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

	adicionarUsuario()
	login()
	realizaPagamento()
	sentry.CaptureMessage("isso é um teste")

	sentry.Flush(time.Second*10)
}

func realizaPagamento() {
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category: "pagamento",
		Message: "Realizando Pagamento",
		Level: sentry.LevelDebug,
	});

	// implementação ...
}

func adicionarUsuario() {
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category: "usuario",
		Message: "Adicionando usuário",
		Level: sentry.LevelDebug,
	});

	// implementação ...
}

func login() {
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category: "login",
		Message: "Logando",
		Level: sentry.LevelDebug,
	});

	// implementação ...
}


