package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vanshaj/kubernetes_practice/webserver/app"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if err := app.Run(); err != nil {
		log.Error().Msg(err.Error())
	}
}
