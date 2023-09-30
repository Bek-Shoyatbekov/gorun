package main

import "os"

func ReadFile(path string) string {
	dat, err := os.ReadFile(path)
	CheckError(err)
	return string(dat)
}
