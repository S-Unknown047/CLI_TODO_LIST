package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `josn:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
