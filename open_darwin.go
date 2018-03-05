package main

import (
	"os/exec"
)

func open(filepath string) {
	if filepath == "" {
		return
	}

	cmd := exec.Command("open", filepath)
	cmd.Run()
}
