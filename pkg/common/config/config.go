package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:DB_URL`
}

//LoadConfig is a function to help us load configurations
func LoadConfig() (c Config, err error) {
	viper.SetConfigName("dev")               
	viper.SetConfigType("env")                 
	viper.AddConfigPath("./pkg/common/envs") 

	viper.AutomaticEnv()

	err = viper.ReadInConfig()  

	if err != nil {                           
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	
	err = viper.Unmarshal(&c)
	return
}
