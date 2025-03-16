package main

import (
	"fmt"
	"os"
)

func Pwd(_ []string) {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("pwd: error getting current directory:", err)
		return
	}

	fmt.Println(dir)
}
