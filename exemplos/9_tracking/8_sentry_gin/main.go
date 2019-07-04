package main

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func main() {
	_ = sentry.Init(sentry.ClientOptions{
		// sentry informa dsn por env. Crie a env SENTRY_DSN=<seu_dsn>
		//Dsn:              "<dsn>",
		Debug:            true,
		AttachStacktrace: true,
	})

	app := gin.Default()

	app.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	app.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	})

	// enviar erro com contexto http
	// - cookie
	// - url
	// ...
	//
	app.GET("/envia-erro", func(ctx *gin.Context) {

		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.CaptureException(errors.New("ocorreu um erro"))
		}

		ctx.Status(http.StatusOK)
	})

	// envia erro simples
	// somente stacktrace
	//
	app.GET("/envia-erro-simples", func(ctx *gin.Context) {

		sentry.CaptureException(errors.New("ocorreu um erro"))

		ctx.Status(http.StatusOK)
	})

	// forcar panic
	//
	app.GET("/forcado", func(ctx *gin.Context) {
		panic("erro forçado")
	})

	_ = app.Run(":8080")
}