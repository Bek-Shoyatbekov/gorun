package main

import ()

var (
	CONFIG_PATH = "./.gorun"
)

func GetConfigs() *Configurations {
	gorunFile := DoesFileExist(CONFIG_PATH)
	if !gorunFile {
		return &Configurations{
			rootDir:  ".",
			exclude:  []string{},
			logError: true,
			delay:    0,
		}
	}

	gorunData := ReadFile(CONFIG_PATH)
	configs := ExtractConfigInfo(gorunData)
	return configs
}