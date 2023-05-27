package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"

	"github.com/vanshaj/kubernetes_practice/webserver/logic"
)

func (s *MyServer) CreateTask(w http.ResponseWriter, r *http.Request) {
	t := &logic.Task{}
	if err := t.Decode(r.Body); err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error": "Incorrect request body"`))
	}
	insert_query := `INSERT INTO TASKS (TASK_NAME, USER_ID) VALUES ($1, $2)`
	if _, err := s.db.Exec(insert_query, t.Details, t.UserId); err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error": "Incorrect data sent"`))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *MyServer) GetTask(w http.ResponseWriter, r *http.Request) {
	t := &logic.Task{}
	taskId := chi.URLParam(r, "id")
	tId, err := strconv.Atoi(taskId)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error": "send correct task id format"`))
	}
	t.Id = tId
	get_task_query := `SELECT TASK_NAME, USER_ID FROM TASKS WHERE ID = $1`
	var name string
	var user_id int
	row := s.db.QueryRow(get_task_query, tId)
	err = row.Scan(&name, &user_id)
	switch err {
	case sql.ErrNoRows:
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error": "No data for this task found"`))
		return
	}
	t.Details = name
	t.UserId = user_id
	if err = t.Encode(w); err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`"error": "Unable to process request"`))
	}
}
