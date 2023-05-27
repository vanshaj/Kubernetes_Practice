package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/vanshaj/kubernetes_practice/webserver/api"
)

func Run() error {
	log.Info().Msg("Starting webserver")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "", "", "taskdb")
	dbCon, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	s := api.NewServer(3000, dbCon)
	return s.StartServer()
}
