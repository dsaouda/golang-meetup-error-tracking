# Error Handler e Error tracking em golang

Para visualizar essa palestra você deve baixar o projeto [Go present tool](https://godoc.org/golang.org/x/tools/cmd/present).

`go get golang.org/x/tools/cmd/present`

execute `present`

Se preferir acesse [https://go-talks.appspot.com/github.com/dsaouda/golang-meetup-error-tracking/slides.slide#1]

# exemplos

1. Para rodar os exemplos você deve criar uma conta no https://sentry.io ou instalar o sentry local https://hub.docker.com/_/sentry/
2. Crie a variável de ambiente SENTRY_DSN com a chave do seu projeto, você encontra no link https://sentry.io/settings/<organization>/projects/<project>/keys/
	
	`export SENTRY_DSN=https://<key>@sentry.io/<project>`

# demo

1. Crie uma conta na total voice (https://totalvoice.com.br)
2. Pegue o access_token em https://api2.totalvoice.com.br/painel/profile.php
3. Crie as variáveis de ambiente TOTAL_VOICE_ACCESS_TOKEN, TOTAL_VOICE_BINA, SENTRY_DSN
	
	`export SENTRY_DSN=https://<key>@sentry.io/<project>`
	
	`export TOTAL_VOICE_ACCESS_TOKEN=<access_token>`
	
	`export TOTAL_VOICE_BINA=<numero_habilitado_para_aparecer_na_bina>`

