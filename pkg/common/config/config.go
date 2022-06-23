package config

import (
	"log"

	"github.com/spf13/viper"
)

// DBConfig struct
type DBConfig struct {
	Port  string 
	DBUrl string
}

//LoadDBConfig helps us to load DB configurations
func LoadDBConfig() (c DBConfig, err error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("../pkg/common/envs")

	err = vp.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}

	c.Port = vp.GetString("port")
	c.DBUrl = vp.GetString("dbUrl")

	return c, err
}
