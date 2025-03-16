package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Cd(args []string) {
	if len(args) == 0 {
		fmt.Println("cd: no argument provided")
		return
	}

	path := args[0]

	if strings.HasPrefix(path, "~") {
		path = strings.TrimPrefix(path, "~")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("cd: error getting home directory:", err)
			return
		}

		path = filepath.Join(homeDir, path)
	}

	if err := os.Chdir(path); err != nil {
		fmt.Println("cd: " + path + ": No such file or directory")
		return
	}
}
