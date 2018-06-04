package dao

import (
	"log"

	. "github.com/marcelocollyer/resume_go_app/config"
	. "github.com/marcelocollyer/resume_go_app/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResumeDAO struct {
}

var db *mgo.Database

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {

	var config = Config{}
	config.Read()

	// Establish a connection to database
	session, err := mgo.Dial(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(config.Database)
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

// Insert a resume into database
func (m *ResumeDAO) Insert(resume Resume) error {
	err := db.C(COLLECTION).Insert(&resume)
	return err
}