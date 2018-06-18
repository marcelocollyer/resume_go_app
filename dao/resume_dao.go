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
	//TODO move this code to dao_util

	var config = Config{}
	config.Read()

	// Establish a connection to database
	session, err := mgo.Dial(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(config.Database)
}

// FindAll - Finds list of resumes
func (m *ResumeDAO) FindAll() ([]Resume, error) {
	var resumes []Resume
	err := db.C(COLLECTION).Find(bson.M{}).All(&resumes)
	return resumes, err
}

// FindByID - Finds a resume by its id
func (m *ResumeDAO) FindByID(id string) (Resume, error) {
	var resume Resume
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&resume)
	return resume, err
}

// Insert - Inserts a resume into database
func (m *ResumeDAO) Insert(resume Resume) error {
	err := db.C(COLLECTION).Insert(&resume)
	return err
}

// Update - Updates a resume into database
func (m *ResumeDAO) Update(resume Resume) error {
	err := db.C(COLLECTION).UpdateId(resume.ID, &resume)
	return err
}

// Delete - Deletes a resume from database
func (m *ResumeDAO) Delete(resume Resume) error {
	err := db.C(COLLECTION).Remove(&resume)
	return err
}

// DeleteByID - Deletes a resume from database by given ID
func (m *ResumeDAO) DeleteByID(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}
