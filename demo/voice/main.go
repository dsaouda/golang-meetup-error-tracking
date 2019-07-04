package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/totalvoice/totalvoice-go"
	"net/http"
	"os"
	"strings"
)

var client *totalvoice.TotalVoice

func init() {
	client = totalvoice.NewTotalVoiceClient(os.Getenv("TOTAL_VOICE_ACCESS_TOKEN"))

	err := sentry.Init(sentry.ClientOptions{
		SampleRate:       1.0,
		Dsn:              os.Getenv("SENTRY_DSN"),
		AttachStacktrace: true,
		Debug:            true,
	})

	if err != nil {
		log.Fatalf("não foi possível iniciar o sentry: %v", err)
	}

}

func main() {
	// Echo instance
	e := echo.New()
	e.Static("/", "public")
	e.File("/", "public/index.html")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))

	// realizar uma ligação
	//
	e.POST("/fazer-chamada", func(c echo.Context) error {

		telefone := c.FormValue("telefone")
		mensagem := c.FormValue("mensagem")

		if strings.TrimSpace(telefone) == "" {
			return c.JSON(400, "telefone deve ser preenchido")
		}

		if len(telefone) < 8 {
			notificarErroTelefone(c, telefone, mensagem)
			return c.JSON(400, "telefone está inválido: quantidade de caracteres")
		}

		if strings.TrimSpace(mensagem) == "" {
			return c.JSON(400, "mensagem deve ser preenchida")
		}

		response, err := fazerLigacaoTTS(telefone, mensagem)

		if err != nil {
			notificarErroTotalVoice(response, err)
			return c.JSON(400, "não foi possível fazer a chamada")
		}

		return c.String(http.StatusOK, "")
	})

	// Start server
	e.Logger.Fatal(e.Start("localhost:1323"))
}

// notificar erro de resposta ou comunicação da total voice
//
func notificarErroTotalVoice(response interface{}, err error) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetExtra("response", response)
		sentry.CaptureException(err)
	})
}

// notificar no sentry que um usuário digitou telefone inválido
//
func notificarErroTelefone(c echo.Context, telefone string, mensagem string) {
	if hub := sentryecho.GetHubFromContext(c); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			scope.SetExtra("telefone_digitado", telefone)
			scope.SetExtra("mensagem", mensagem)
			hub.CaptureMessage(fmt.Sprintf("alguém digitou esse telefone (%v) e não está correto", telefone))
		})
	}
}

// realizar uma ligação TTS
//
func fazerLigacaoTTS(telefone, mensagem string) (interface{}, error) {

	opcoes := map[string]interface{}{
		"velocidade": 1,
		"tipo_voz":   "br-Vitoria",
		"bina":       os.Getenv("TOTAL_VOICE_BINA"),
	}

	response, err := client.TTS.Enviar(telefone, mensagem, opcoes)

	if err != nil {
		return nil, errors.WrapPrefix(err, "erro na comunicação com total voice", 0)
	}

	if response.Status != 200 {
		return response, errors.New("os dados enviados não foram validados")
	}

	return nil, nil
}
