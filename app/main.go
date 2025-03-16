package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		name, args := parse(input)

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
