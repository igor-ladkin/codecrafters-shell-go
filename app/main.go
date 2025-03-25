package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("main: error reading input")
			os.Exit(1)
		}

		name, args, io := nameAndArgs(input)

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
