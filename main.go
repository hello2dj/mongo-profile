package main

import (
	"mongo-profile/db"
	"mongo-profile/model"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

func main() {

	_, err := db.Connect("localhost:27017", "test")
	println("err is ", err)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		iter := model.GetProfileListByType(model.ByDay)
		var item model.AggProfileLog
		db.DefaultDB.C("system.profile").Find(bson.M{}).Select(bson.M{"_id": 1, "ts": 1}).One(&item)

		println("data", iter.Next(&item))
		if iter.Err() != nil {
			println("error is ", iter.Err().Error())
		}
		println("hello2", item.AvgTime, item.Ts, item.Ds)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:20001")
}
