package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

var builtins = []string{"exit", "echo", "type"}

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
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
}

func handleType(args []string) {
	if len(args) == 0 {
		panic("No command provided")
	}

	name := args[0]

	switch {
	case slices.Contains(builtins, name):
		fmt.Fprintln(os.Stdout, name+" is a shell builtin")
	default:
		fmt.Fprintln(os.Stdout, name+": not found")
	}
}

func handleCommand(command string) {
	parts := strings.Fields(strings.TrimSpace(command))

	if len(parts) == 0 {
		panic("No command provided")
	}

	name, args := parts[0], parts[1:]

	switch name {
	case "exit":
		handleExit(args)
	case "echo":
		handleEcho(args)
	case "type":
		handleType(args)
	default:
		fmt.Println(name + ": command not found")
	}
}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	handleCommand(command)

	main()
}
