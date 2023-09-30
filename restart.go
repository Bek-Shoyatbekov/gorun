package main

import (
	"log"
	"os/exec"
	"runtime"
)

const OS = runtime.GOOS

func Restart() {
	hasGoMod := DoesFileExist("./go.mod")
	var cmd *exec.Cmd
	if hasGoMod == false {
		cmd = exec.Command("go mod init gorun.com")
		cmd.Run()
	}
	cmd = exec.Command("go run .")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		Restart()
	}
}
