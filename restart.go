package main

import (
	"os"
	"os/exec"
	"runtime"
)

const OS = runtime.GOOS

var (
	PROCESS_ID = os.Getgid()
	cmd        *exec.Cmd
)

func Start() {
	cmd = exec.Command("go", "run", ".")
	cmd.Start()
}

func Restart() {
	cmd.Process.Kill()
	hasGoMod := DoesFileExist("./go.mod")
	if !hasGoMod {
		cmd = exec.Command("go", "mod", "init", "gorun")
		err := cmd.Start()
		CheckError(err)
	}
	go Start()
}
