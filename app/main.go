package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func handleExit(args []string) {
	code, _ := strconv.Atoi(args[0])
	os.Exit(code)
}

func handleEcho(args []string) {
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
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
