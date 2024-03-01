package model

import "time"

type employee struct {
	ID         string    `json:"id"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Department string    `json:"department"`
	Age        string    `json:"age"`
	Create_at  time.Time `json:"create_at"`
}
