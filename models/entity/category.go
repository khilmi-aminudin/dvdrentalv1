package entity

import "time"

type Category struct {
	CategoryId int64     `json:"category_id"`
	Name       string    `json:"name"`
	LastUpdate time.Time `json:"last_update"`
}
