package main

import (
	"api/basic-rest-api/pkg/common/db"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbURL := viper.Get("DB_URL").(string)

	fmt.Println(port, dbURL)

	r := gin.Default()

	db.Init(dbURL)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbURL,
		})
	})
	r.Run(port)
}
