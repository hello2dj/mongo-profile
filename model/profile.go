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

type QueryType int8

const (
	ByDay = iota
	ByMonth
	ByYear
)

type AggProfileLog struct {
	Ts      int     `bson:"ts"`
	Ds      string  `bson:"ds"`
	AvgTime float32 `bson:"avgTime"`
}

// var addEle = [1]bson.M{bson.M{"$ts": 28800000}}

var projToAdd8Hour = bson.M{
	"$project": bson.M{
		"ts": 1,
		// "ts": bson.M{
		// 	"$add": addEle,
		// },
		"millis": 1,
	},
}

func GetProfileListByType(t QueryType) *mgo.Iter {
	var pipeline [4]bson.M

	// pipeline[0] = projToAdd8Hour
	switch t {
	case ByDay:
		var projToGetDate = bson.M{
			"$project": bson.M{
				"ts": 1,
				"ds": bson.M{ // date string
					"$dateToString": bson.M{
						"format": "%Y-%m-%d",
						"date":   "$ts",
					},
				},
				"millis": 1,
			},
		}
		pipeline[0] = projToGetDate
		pipeline[1] = bson.M{"$group": bson.M{"_id": bson.M{
			"ts": bson.M{"$dayOfMonth": "$ts"},
			"ds": "$ds",
		},
			"avgTime": bson.M{
				"$avg": "$millis",
			},
		}}
	case ByMonth:
		var projToGetDate = bson.M{
			"$project": bson.M{
				"ts": 1,
				"ds": bson.M{ // date string
					"$dateToString": bson.M{
						"format": "%Y-%m",
						"date":   "$ts",
					},
				},
				"millis": 1,
			},
		}
		pipeline[0] = projToGetDate
		pipeline[1] = bson.M{"$group": bson.M{"_id": bson.M{
			"ts": bson.M{"$month": "$ts"},
			"ds": "$ds",
		},
			"avgTime": bson.M{
				"$avg": "$millis",
			},
		}}
	case ByYear:
		var projToGetDate = bson.M{
			"$project": bson.M{
				"ts": 1,
				"ds": bson.M{ // date string
					"$dateToString": bson.M{
						"format": "%Y",
						"date":   "$ts",
					},
				},
				"millis": 1,
			},
		}
		pipeline[0] = projToGetDate
		pipeline[1] = bson.M{"$group": bson.M{"_id": bson.M{
			"ts": bson.M{"$year": "$ts"},
			"ds": "$ds",
		},
			"avgTime": bson.M{
				"$avg": "$millis",
			},
		}}
	default:
		var projToGetDate = bson.M{
			"$project": bson.M{
				"ts": 1,
				"ds": bson.M{ // date string
					"$dateToString": bson.M{
						"format": "%Y-%m-%d",
						"date":   "$ts",
					},
				},
				"millis": 1,
			},
		}
		pipeline[0] = projToGetDate
		pipeline[1] = bson.M{"$group": bson.M{"_id": bson.M{
			"ts": bson.M{"$dayOfMonth": "$ts"},
			"ds": "$ds",
		},
			"avgTime": bson.M{
				"$avg": "$millis",
			},
		}}
	}
	pipeline[2] = bson.M{"$project": bson.M{"ts": "$_id.ts", "ds": "$_id.ds", "avgTime": 1}}
	pipeline[3] = bson.M{"$project": bson.M{"_id": 0, "ts": 1, "ds": 1, "avgTime": 1}}
	return db.DefaultDB.C("system.profile").Pipe(pipeline).Iter()
}
