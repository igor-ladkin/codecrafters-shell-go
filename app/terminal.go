package main

import (
	"os"
	"strings"

	"golang.org/x/term"
)

func setupTerminal() *term.Terminal {
	terminal := term.NewTerminal(os.Stdin, "\r$ ")
	terminal.AutoCompleteCallback = autocomplete

	return terminal
}

func autocomplete(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key != 9 {
		return "", 0, false
	}

	if strings.HasPrefix("echo", line) {
		return "echo ", len("echo "), true
	}

	if strings.HasPrefix("exit", line) {
		return "exit ", len("exit "), true
	}

	return "", 0, false
}
