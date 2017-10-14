package controller

import (
	"mongo-profile/model"
	"mongo-profile/service"

	"github.com/gin-gonic/gin"
)

func GetProfileStats(ctx *gin.Context) {
	iter := model.Aggregate(service.GetProfileListByType(service.ByDay))
	var item model.AggProfileLog
	// db.DefaultDB.C("system.profile").Find(bson.M{}).Select(bson.M{"_id": 1, "ts": 1}).One(&item)

	iter.Next(&item)

	if iter.Err() != nil {
		println("error is ", iter.Err().Error())
	}

	ctx.JSON(200, gin.H{
		"message": "pong",
		"ts":      item.Ts,
		"avgTime": item.AvgTime,
	})
}
