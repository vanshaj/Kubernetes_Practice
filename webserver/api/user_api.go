package api

import (
	"net/http"

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
