package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

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

	switch name {
	case "exit":
		Exit(args)
	case "echo":
		Echo(args)
	case "type":
		Type(args)
	case "pwd":
		Pwd(args)
	case "cd":
		Cd(args)
	default:
		Exec(name, args)
	}

	main()
}
