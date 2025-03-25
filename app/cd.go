package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Cd(args []string, io IO) {
	if len(args) == 0 {
		fmt.Fprintln(io.Error, "cd: no argument provided")
		return
	}

	path := args[0]

	if strings.HasPrefix(path, "~") {
		path = strings.TrimPrefix(path, "~")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(io.Error, "cd: error getting home directory")
			return
		}

		path = filepath.Join(homeDir, path)
	}

	if err := os.Chdir(path); err != nil {
		fmt.Fprintln(io.Error, "cd: "+path+": no such file or directory")
		return
	}
}
