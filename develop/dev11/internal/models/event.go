package models

import "time"

type Event struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
