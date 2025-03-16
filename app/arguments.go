package main

import (
	"bufio"
	"strings"
	"unicode"
)

func parse(input string) (string, []string) {
	var name string
	var args []string

	parts := strings.SplitAfterN(strings.TrimSpace(input), " ", 2)

	if len(parts) == 0 {
		panic("No command provided")
	}

	if len(parts) > 1 {
		args = parseArgs(parts[1])
	}

	name = parseName(parts[0])

	return name, args
}

func parseName(str string) string {
	return strings.TrimSpace(str)
}

func parseArgs(input string) []string {
	var args []string
	var currentArg strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)

	isQuoted := false

	for scanner.Scan() {
		char := scanner.Text()[0]
		if char == '\'' {
			if isQuoted {
				isQuoted = false
			} else {
				isQuoted = true
			}
			continue
		}

		if unicode.IsSpace(rune(char)) && !isQuoted {
			if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
			continue
		}

		currentArg.WriteByte(char)
	}

	if currentArg.Len() > 0 {
		args = append(args, strings.TrimSpace(currentArg.String()))
	}

	return args
}
