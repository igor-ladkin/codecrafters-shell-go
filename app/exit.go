package main

import (
	"fmt"
	"os"
	"strconv"
)

func Exit(args []string) {
	code, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("exit: Invalid exit code")
		return
	}

	if code < 0 || code > 255 {
		fmt.Println("exit: Invalid exit code")
		return
	}

	os.Exit(code)
}
