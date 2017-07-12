package main

import "os/exec"

func execCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	return cmd.Run()
}
