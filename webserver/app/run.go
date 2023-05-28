package app

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/vanshaj/kubernetes_practice/webserver/api"
)

func Run() error {
	log.Info().Msg("Starting webserver")
	// get env host, port, user, password, dbname
	host := os.Getenv("DB_HOST")
	port, bool := os.LookupEnv("DB_PORT")
	if !bool {
		port = "5432"
	}
	portNumber, _ := strconv.Atoi(port)
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portNumber, user, password, dbname)
	dbCon, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	s := api.NewServer(3000, dbCon)
	return s.StartServer()
}
