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

func handleExit(args []string) {
	code, err := strconv.Atoi(args[0])

	if err != nil {
		panic("Invalid exit code")
	}

	if code < 0 || code > 255 {
		panic("Invalid exit code")
	}

	os.Exit(code)
}

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func handleType(args []string) {
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

func handlePwd(_ []string) {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("pwd: error getting current directory:", err)
		return
	}

	fmt.Println(dir)
}

func handleCd(args []string) {
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

func handleBuiltin(name string, args []string) {
	switch name {
	case "exit":
		handleExit(args)
	case "echo":
		handleEcho(args)
	case "type":
		handleType(args)
	case "pwd":
		handlePwd(args)
	case "cd":
		handleCd(args)
	}
}

func handleExecutable(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func handleCommand(command string) {
	parts := strings.Fields(strings.TrimSpace(command))

	if len(parts) == 0 {
		panic("No command provided")
	}

	name, args := parts[0], parts[1:]

	if isBuiltin(name) {
		handleBuiltin(name, args)
		return
	}

	if _, ok := isExecutable(name); ok {
		handleExecutable(name, args)
		return
	}

	fmt.Println(name + ": command not found")
}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	handleCommand(command)

	main()
}
