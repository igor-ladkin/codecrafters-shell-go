package main

import (
	"fmt"
	"os/exec"
)

func Exec(name string, args []string, io IO) error {
	if _, err := exec.LookPath(name); err != nil {
		fmt.Fprintln(io.Error, name+": command not found")
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = io.Input
	cmd.Stdout = io.Output
	cmd.Stderr = io.Error

	return cmd.Run()
}
