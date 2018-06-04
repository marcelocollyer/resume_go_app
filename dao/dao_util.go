package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type DAOUtil struct {
	Server   string
	Database string
	db       *mgo.Database
}

// Establish a connection to database
func (m *DAOUtil) Connect() *mgo.Database {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(m.Database)
}
