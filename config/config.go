package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// declare config globally
var config *viper.Viper

func Init(env string) {
	// initialise err variable with type error
	var err error
	config = viper.New()

	// set file format
	config.SetConfigType("yaml")

	// set the config file name
	config.SetConfigName(env)

	// set config file path
	config.AddConfigPath("config/")
	
	// find file folder
	err = config.ReadInConfig()
	
	// if file does not exist
	if err != nil {
		panic(fmt.Errorf("error file not found %w", err))
	}
}

func GetConfig(env string) *viper.Viper{
	// initialize environment files
	Init(env)

	return config
}