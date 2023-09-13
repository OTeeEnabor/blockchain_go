package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func getConfigFile() string {
	// set file name
	path, _ := os.Executable()
	// get file path
	filePath := filepath.Dir(path)
	// store  template folder path into variable (CLOUD server)
	configFile := fmt.Sprintf("%v/config/production.yaml",filePath)
	
	return configFile
}


// var EnVar = GetConfig("development")
// set file name
var configFilePath = getConfigFile()
var EnVar = GetConfig(configFilePath)