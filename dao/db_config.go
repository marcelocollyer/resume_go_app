package dao

import (
	"log"

	. "github.com/marcelocollyer/resume_go_app/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResumeDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "resumes"
)

// Establish a connection to database
func (m *ResumeDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of resumes
func (m *ResumeDAO) FindAll() ([]Resume, error) {
	var resumes []Resume
	err := db.C(COLLECTION).Find(bson.M{}).All(&resumes)
	return resumes, err
}

// Find a resume by its id
func (m *ResumeDAO) FindById(id string) (Resume, error) {
	var resume Resume
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&resume)
	return resume, err
}

// Insert a movie into database
func (m *ResumeDAO) Insert(resume Resume) error {
	err := db.C(COLLECTION).Insert(&resume)
	return err
}
