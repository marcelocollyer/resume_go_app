package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Resume struct (Model)
type Resume struct {
	ID                bson.ObjectId   `bson:"_id" json:"id"`
	Name              string          `json:"name"`
	ProfileImgURL     string          `json:"profile_img_url"`
	Email             string          `json:"email"`
	Phone             string          `json:"phone"`
	Skype             string          `json:"skype"`
	Website           string          `json:"website"`
	Linkedin          string          `json:"linkedin"`
	Github            string          `json:"github"`
	Twitter           string          `json:"twitter"`
	RolesDesc         string          `json:"roles_desc"`
	CarrerDesc        string          `json:"carrer_desc"`
	ExperienceList    []Experience    `json:"experiences"`
	EducationList     []Education     `json:"educations"`
	LanguageList      []Language      `json:"languages"`
	CertificationList []Certification `json:"certifications"`
	InterestList      []Interest      `json:"interests"`
	SkillList         []Skill         `json:"skills"`
}

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

// Certification struct model
type Certification struct {
	Description string `json:"description"`
	Company     string `json:"company"`
}

// Interest struct model
type Interest struct {
	Description string `json:"description"`
}

// Skill struct model
type Skill struct {
	Description string  `json:"description"`
	Level       string  `json:"level"`
	Skill       []Skill `json:"children"`
}
