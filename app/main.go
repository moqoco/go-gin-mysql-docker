package main

import (
	"api/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB connection gom
	model.Conn()

	r := gin.Default()
	r.GET("/sample", runSample)
	r.Run(":8000")
}


func runSample(c *gin.Context) {
	sample := model.Sample{}
	sample.First()
	
	c.JSON(200, sample)
}