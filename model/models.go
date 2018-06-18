package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Experience struct model
type Experience struct {
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

// Education struct model
type Education struct {
	Title     string    `json:"title"`
	Course    string    `json:"course"`
	Location  string    `json:"location"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Language struct model
type Language struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type Interest struct {
	Description string `json:"description"`
}

type Skill struct {
	Description string `json:"description"`
	Level       string `json:"level"`
}

// Resume struct (Model)
type Resume struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Name           string        `json:"name"`
	ProfileImgURL  string        `json:"profile_img_url"`
	Email          string        `json:"email"`
	Phone          string        `json:"phone"`
	Website        string        `json:"website"`
	Linkedin       string        `json:"linkedin"`
	Github         string        `json:"github"`
	Twitter        string        `json:"twitter"`
	RolesDesc      string        `json:"roles_desc"`
	CarrerDesc     string        `json:"carrer_desc"`
	ExperienceList []Experience  `json:"experiences"`
	EducationList  []Education   `json:"educations"`
	LanguageList   []Language    `json:"languages"`
	InterestList   []Interest    `json:"interests"`
	SkillList      []Skill       `json:"skills"`
}
