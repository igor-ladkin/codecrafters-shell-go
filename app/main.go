package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	terminal := setupTerminal()

	for {
		input, err := terminal.ReadLine()

		if err != nil {
			fmt.Println("main: error reading input")
			os.Exit(1)
		}

		name, args, io := parseArguments(input)

		switch name {
		case "exit":
			Exit(args, io)
		case "echo":
			Echo(args, io)
		case "type":
			Type(args, io)
		case "pwd":
			Pwd(args, io)
		case "cd":
			Cd(args, io)
		default:
			Exec(name, args, io)
		}
	}
}
