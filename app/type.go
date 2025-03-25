package main

import (
	"fmt"
	"os/exec"
	"slices"
)

var builtins = []string{
	"exit",
	"echo",
	"type",
	"pwd",
	"cd",
}

func isBuiltin(name string) bool {
	return slices.Contains(builtins, name)
}

func isExecutable(name string) (string, bool) {
	path, err := exec.LookPath(name)
	return path, err == nil
}

func Type(args []string, io IO) {
	if len(args) == 0 {
		fmt.Fprintln(io.Error, "type: no command provided")
		return
	}

	name := args[0]

	if isBuiltin(name) {
		fmt.Fprintln(io.Output, name, "is a shell builtin")
		return
	}

	if path, ok := isExecutable(name); ok {
		fmt.Fprintln(io.Output, name, "is", path)
		return
	}

	fmt.Fprintln(io.Error, name+": not found")
}
