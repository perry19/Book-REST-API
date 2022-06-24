package main

import (
	"api/basic-rest-api/pkg/books"
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

	r := gin.Default()
	h := db.Init(dbURL)

	books.RegisterRoutes(r, h)

	r.Run(port)
}
