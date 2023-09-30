package main

import (
	"log"
	"strconv"
	"strings"
)

func ExtractConfigInfo(gorunInfo string) *Configurations {
	validInfo := checkRequiredInfo(gorunInfo)
	if !validInfo {
		log.Fatal("Invalid .gorun configurations")
	}
	info := strings.Split(gorunInfo, "\n")
	info = removeEmpty(info)
	exclude := strings.Split(strings.Split(info[0], ":")[1], ",")
	root := strings.Split(info[1], ":")[1]
	logError, _ := strconv.ParseBool(strings.Split(info[2], ":")[1])
	delay, _ := strconv.ParseInt(strings.Split(info[3], ":")[1], 2, 64)
	return &Configurations{
		rootDir:  root,
		exclude:  exclude,
		logError: logError,
		delay:    int32(delay),
	}
}

func checkRequiredInfo(info string) bool {
	ok := true
	requireds := []string{"rootDir", "excludes", "logErrors"}
	for _, v := range requireds {
		if !strings.Contains(info, v) {
			ok = false
		}
	}
	return ok
}

func removeEmpty(arr []string) []string {
	emp := ""
	var newArr []string
	for _, e := range arr {
		if e != emp {
			newArr = append(newArr, e)
		}
	}
	return newArr
}
