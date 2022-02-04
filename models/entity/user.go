package entity

import "time"

type Users struct {
	UserId     int64     `json:"user_id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Passowrd   string    `json:"password"`
	Tokens     string    `json:"tokens"`
	LastUpdate time.Time `json:"last_update"`
}
