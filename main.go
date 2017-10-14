package main

import (
	"mongo-profile/controller"
	"mongo-profile/db"

	"github.com/gin-gonic/gin"
)

func main() {

	_, err := db.Connect("localhost:27017", "test")
	println("err is ", err)

	r := gin.Default()

	r.GET("/", controller.GetProfileStats)
	r.Run("localhost:20001")
}
