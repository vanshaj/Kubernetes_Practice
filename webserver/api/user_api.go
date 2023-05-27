package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"

	"github.com/vanshaj/kubernetes_practice/webserver/logic"
)

func (s *MyServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := &logic.User{}
	if err := u.Decode(r.Body); err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error": "Invalid request body"`))
		return
	}

	create_query := `INSERT INTO USERS(name, dept) values ($1, $2)`
	s.db.QueryRow(create_query, u.Name, u.Dept)
	w.WriteHeader(http.StatusCreated)
}

func (s *MyServer) GetUser(w http.ResponseWriter, r *http.Request) {
	u := &logic.User{}
	userId := chi.URLParam(r, "id")
	uId, err := strconv.Atoi(userId)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("User id sent is incorrect '%s'", userId))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error": "Incorrect format of userId"`))
		return
	}
	u.Id = uId
	log.Debug().Msg(fmt.Sprintf("UserId fetched from the url param is '%d'", uId))
	get_user_query := `SELECT NAME, DEPT FROM USERS WHERE Id = $1;`
	row := s.db.QueryRow(get_user_query, uId)
	var name string
	var dept string
	err = row.Scan(&name, &dept)
	switch err {
	case sql.ErrNoRows:
		log.Error().Msg("No rows were found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error": "Incorrect id sent"`))
		return
	}
	u.Name = name
	u.Dept = dept
	if err := u.Encode(w); err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`"error": "Unable to process request"`))
		return
	}
}
