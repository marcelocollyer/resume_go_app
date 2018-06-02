package model

import "gopkg.in/mgo.v2/bson"

// Resume struct (Model)
type Resume struct {
	ID          bson.ObjectId `json:"id"`
	Name        string        `json:"name"`
	CareerDesc  string        `json:"carrer_desc"`
	RolesDesc   string        `json:"roles_desc"`
	Experiences []Experience  `json:"experiences"`
}
