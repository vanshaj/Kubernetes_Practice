package logic

import (
	"encoding/json"
	"io"
)

type Task struct {
	Id      int    `json:"id"`
	Details string `json:"details"`
	UserId  int    `json:"userId"`
}

func NewTask(id int, details string, userId int) *Task {
	return &Task{
		Id:      id,
		Details: details,
		UserId:  userId,
	}
}

func (u *Task) Encode(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(u)
}

func (u *Task) Decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(u)
}
