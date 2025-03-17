package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec(name string, args []string) error {
	if _, err := exec.LookPath(name); err != nil {
		fmt.Println("exec: " + name + ": command not found")
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
