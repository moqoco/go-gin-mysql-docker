package main

import (
	"log"
	"api/model"
	"github.com/gin-gonic/gin"
)


type Sample struct {
	ID   uint	`json:"id"`
	Txt  string `json:"txt"`
}

func main() {
	db := model.Conn()
	defer db.Close()

	results, err := db.Query("SELECT * FROM sample")
	if err != nil {
		log.Fatal(err)
	}
	defer results.Close()
	
	for results.Next() {
		var sample model.Sample

		err = results.Scan(&sample.ID, &sample.Txt)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(sample.Txt)
	}


}

func run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000")
}

