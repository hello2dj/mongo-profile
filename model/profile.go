package model

import (
	"mongo-profile/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type OpEnum string

const (
	Command  = "command"
	Count    = "count"
	Distinct = "distinct"
	Insert   = "insert"
	Query    = "query"
	Remove   = "remove"
	Update   = "update"
)

type AggProfileLog struct {
	Ts      int     `bson:"ts"`
	Ds      string  `bson:"ds"`
	AvgTime float32 `bson:"avgTime"`
}

func Aggregate(pipeline []bson.M) *mgo.Iter {
	return db.DefaultDB.C("system.profile").Pipe(pipeline).Iter()
}

func FindOne(query bson.M) *mgo.Query {
	return db.DefaultDB.C("system.profile").Find(query)
}

func FindById(id string) *mgo.Query {
	return db.DefaultDB.C("system.profile").Find(bson.M{"_id": id})
}
