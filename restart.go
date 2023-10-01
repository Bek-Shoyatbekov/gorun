package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

const OS = runtime.GOOS

func Restart() {
	hasGoMod := DoesFileExist("./go.mod")
	app := "go"
	mod := "mod"
	run := "run"
	init := "init"
	arg0 := "."
	arg1 := "gorun.com"
	var cmd *exec.Cmd
	if !hasGoMod {
		cmd = exec.Command(app, mod, init, arg1)
		err := cmd.Run()
		CheckError(err)
	}
	var err error
	cmd = exec.Command(app, run, arg0)
	stdout, _ := cmd.Output()
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		Restart()
		fmt.Println(string(stdout))
	}
}
