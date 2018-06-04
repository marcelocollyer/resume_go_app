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

// Connect - Establishes a connection to database TODO: adapt code to use this function instead
func (m *DAOUtil) Connect() *mgo.Database {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(m.Database)
}
