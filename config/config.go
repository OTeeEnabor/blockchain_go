package config

import (
	"fmt"
	"os"
	"path/filepath"

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
	// set file name
	path, _ := os.Executable()
	// // get file path
	filePath := filepath.Dir(path)
	// store  template folder path into variable (CLOUD server)
	configFolder := fmt.Sprintf("%v/config/",filePath)

	// set config folder path
	config.AddConfigPath(configFolder)
	
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