package meuerro

import (
	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"
)

func init() {
	_ = sentry.Init(sentry.ClientOptions{
		Debug:            	false,
		SampleRate: 		1.0,
		//Dsn:              "<dsn>", // sentry informa dsn por env
		AttachStacktrace: 	true,
	})
}

func ClientNaoEncontradoError() error {
	return errors.New("Cliente não encontrado")
}

func ArgumentWithError(e error) error {
	return Wrap(e, ArgumentError().Error())
}

func ArgumentError() error {
	return errors.New("argument error")
}

func NotFoundError() error {
	return errors.New("not found error")
}

func Wrap(e error, message string) error {
	return errors.WrapPrefix(e, message, 0)
}

func Notify(e error) {
	sentry.CaptureException(e)
}