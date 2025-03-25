package main

import (
	"fmt"
	"os"
)

func Pwd(_ []string, io IO) {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Fprintln(io.Error, "pwd: error getting current directory")
		return
	}

	fmt.Fprintln(io.Output, dir)
}
