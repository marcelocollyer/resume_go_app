package model

import "time"

// Experience struct model
type Experience struct {
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
