package models

import (
	"time"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validat:"required"`
}

type Message struct {
	Sender string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
