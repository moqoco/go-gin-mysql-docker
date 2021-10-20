package main

import (
	"log"
	
	"api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// read env
	err := godotenv.Load()

	if err != nil {
		log.Print("Not Found .env file.")
	}

	// DB connection gorm
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