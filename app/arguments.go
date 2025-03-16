package main

import (
	"bufio"
	"strings"
	"unicode"
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

		if char == '\\' && !isEscaped {
			isEscaped = true
			continue
		}

		if char == '"' && !isEscaped && !isSingleQuoted {
			isDoubleQuoted = !isDoubleQuoted
			continue
		}

		if char == '\'' && !isEscaped && !isDoubleQuoted {
			isSingleQuoted = !isSingleQuoted
			continue
		}

		if char == ' ' && isEscaped && (isSingleQuoted || isDoubleQuoted) {
			isEscaped = false
			currentArg.WriteByte('\\')
			currentArg.WriteByte(char)
			continue
		}

		if char == '\\' && isEscaped && (isSingleQuoted || isDoubleQuoted) {
			isEscaped = false
			currentArg.WriteByte(char)
			continue
		}

		if char == '"' && isEscaped && isSingleQuoted {
			isEscaped = false
			currentArg.WriteByte('\\')
			currentArg.WriteByte(char)
			continue
		}

		if char == '\'' && isEscaped && isDoubleQuoted {
			isEscaped = false
			currentArg.WriteByte('\\')
			currentArg.WriteByte(char)
			continue
		}

		if char == ' ' && isEscaped {
			isEscaped = false
			currentArg.WriteByte(char)
			continue
		}

		if char == '\'' && isEscaped {
			isEscaped = false
			currentArg.WriteByte(char)
			continue
		}

		if char == '"' && isEscaped {
			isEscaped = false
			currentArg.WriteByte(char)
			continue
		}

		if isEscaped && !isSingleQuoted && !isDoubleQuoted {
			isEscaped = false
			currentArg.WriteByte(char)
			continue
		}

		if isEscaped {
			isEscaped = false
			currentArg.WriteByte('\\')
			currentArg.WriteByte(char)
			continue
		}

		if unicode.IsSpace(rune(char)) && !isSingleQuoted && !isDoubleQuoted && !isEscaped {
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
