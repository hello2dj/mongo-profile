package db

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	// b "gopkg.in/mgo.v2/bson"
)

// Session is the gloabl session
var (
	Session   *mgo.Session
	DefaultDB *mgo.Database
)

// Connect is to connect to the mongo
func Connect(url string, db string) (*mgo.Session, error) {
	var err error
	Session, err = mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	if db == "" {
		db = "system.profile"
	}
	DefaultDB = Session.DB(db)
	return Session, nil
}

func GetDb(db string) (*mgo.Database, error) {
	if Session == nil {
		return nil, errors.New("No dial the db")
	}
	return Session.DB(db), nil
}
