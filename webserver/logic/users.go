package logic

import (
	"encoding/json"
	"io"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Dept string `json:"dept"`
}

func NewUser(id int, name, dept string) *User {
	return &User{
		Id:   id,
		Name: name,
		Dept: dept,
	}
}

func (u *User) Encode(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(u)
}

func (u *User) Decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(u)
}
