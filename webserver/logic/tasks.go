package logic

import (
	"encoding/json"
	"io"
)

type Task struct {
	Id      int
	Details string
	UserId  int
}

func NewTask(id int, details string, userId int) *Task {
	return &Task{
		Id:      id,
		Details: details,
		UserId:  userId,
	}
}

func (t *Task) Encode(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(t)

}
