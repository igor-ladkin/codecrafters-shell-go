package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func isBuiltin(name string) bool {
	builtins := []string{"exit", "echo", "type", "pwd", "cd"}
	return slices.Contains(builtins, name)
}

func isExecutable(name string) (string, bool) {
	path, err := exec.LookPath(name)
	return path, err == nil
}

func exit(args []string) {
	code, err := strconv.Atoi(args[0])

	if err != nil {
		panic("Invalid exit code")
	}

	if code < 0 || code > 255 {
		panic("Invalid exit code")
	}

	os.Exit(code)
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func _type(args []string) {
	if len(args) == 0 {
		panic("No command provided")
	}

	name := args[0]

	if isBuiltin(name) {
		fmt.Println(name, "is a shell builtin")
		return
	}

	if path, ok := isExecutable(name); ok {
		fmt.Println(name, "is", path)
		return
	}

	fmt.Println(name + ": not found")
}

func pwd(_ []string) {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("pwd: error getting current directory:", err)
		return
	}

	fmt.Println(dir)
}

func cd(args []string) {
	if len(args) == 0 {
		fmt.Println("cd: no argument provided")
		return
	}

	path := args[0]

	if err := os.Chdir(path); err != nil {
		fmt.Println("cd: " + path + ": No such file or directory")
		return
	}
}

func executeBuiltin(name string, args []string) {
	switch name {
	case "exit":
		exit(args)
	case "echo":
		echo(args)
	case "type":
		_type(args)
	case "pwd":
		pwd(args)
	case "cd":
		cd(args)
	}
}

func executeExternal(name string, args []string) error {
	if _, ok := isExecutable(name); !ok {
		fmt.Println(name + ": command not found")
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	command := strings.Fields(strings.TrimSpace(input))

	if len(command) == 0 {
		panic("No command provided")
	}

	name, args := command[0], command[1:]

	if isBuiltin(name) {
		executeBuiltin(name, args)
	} else {
		executeExternal(name, args)
	}

	main()
}
