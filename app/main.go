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

		name, args := nameAndArgs(input)

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
	}
}
