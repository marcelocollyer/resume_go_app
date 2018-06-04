package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Experience struct model
type Experience struct {
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Resume struct (Model)
type Resume struct {
	ID          bson.ObjectId `json:"id"`
	Name        string        `json:"name"`
	CareerDesc  string        `json:"carrer_desc"`
	RolesDesc   string        `json:"roles_desc"`
	Experiences []Experience  `json:"experiences"`
}
