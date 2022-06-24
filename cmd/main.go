package main

import (
	"api/basic-rest-api/pkg/common/config"
	"api/basic-rest-api/pkg/common/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadDBConfig()

	if err != nil {
		log.Fatalln(err)
	}

	port := config.Port
	dbURL := config.DBUrl
	message := "Hello World"
	r := gin.Default()

	db.Init(dbURL)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbURL,
			"message":message,
		})
	})
	r.Run(port)
}
