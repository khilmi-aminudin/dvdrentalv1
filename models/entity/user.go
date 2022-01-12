package entity

import "time"

type Users struct {
	UserId     int64     `json:"user_id"`
	Username   string    `json:"username"`
	Passowrd   string    `json:"password"`
	LastUpdate time.Time `json:"last_update"`
}
