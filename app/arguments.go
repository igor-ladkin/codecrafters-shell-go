package main

import (
	"bufio"
	"strings"
)

func nameAndArgs(input string) (string, []string) {
	var name string
	var args []string

	if len(input) == 0 {
		panic("No command provided")
	}

	parts := split(input)

	name = parts[0]
	if len(parts) > 1 {
		args = parts[1:]
	}

	return name, args
}

func split(input string) []string {
	var args []string
	var currentArg strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)

	isEscaped := false
	isDoubleQuoted := false
	isSingleQuoted := false

	for scanner.Scan() {
		char := scanner.Text()[0]

		switch {
		case isEscaped:
			isEscaped = false

			if isSingleQuoted || isDoubleQuoted {
				if char == ' ' ||
					(char == '"' && isSingleQuoted) ||
					(char == '\'' && isDoubleQuoted) ||
					(char != '\\' && char != '\'' && char != '"') {
					currentArg.WriteByte('\\')
				}
			}

			currentArg.WriteByte(char)
		case char == '\\':
			isEscaped = true
		case char == '"' && !isSingleQuoted:
			isDoubleQuoted = !isDoubleQuoted
		case char == '\'' && !isDoubleQuoted:
			isSingleQuoted = !isSingleQuoted
		case char == ' ' && !isSingleQuoted && !isDoubleQuoted:
			if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
		default:
			currentArg.WriteByte(char)
		}
	}

	if currentArg.Len() > 0 {
		args = append(args, strings.TrimSpace(currentArg.String()))
	}

	return args
}
