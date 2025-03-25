package main

import (
	"fmt"
	"os"
	"strconv"
)

func Exit(args []string, io IO) {
	code, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Fprintln(io.Error, "exit: invalid exit code")
		return
	}

	if code < 0 || code > 255 {
		fmt.Fprintln(io.Error, "exit: invalid exit code")
		return
	}

	os.Exit(code)
}
