package config

// "os"
// "path/filepath"

// func getConfigFolder() string {
// 	// set file name
// 	path, _ := os.Executable()
// 	// get file path
// 	filePath := filepath.Dir(path)
// 	// store  template folder path into variable (CLOUD server)
// 	configFile := fmt.Sprintf("%v/config/",filePath)

// 	return configFile
// }

// var EnVar = GetConfig("development")
// // set file name
// var configFilePath = getConfigFolder()
var EnVar = GetConfig("production")