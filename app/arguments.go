package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type IO struct {
	Input  io.Reader
	Output io.Writer
	Error  io.Writer
}

func NewIO(input io.Reader, output io.Writer, error io.Writer) IO {
	return IO{
		Input:  input,
		Output: output,
		Error:  error,
	}
}

func NewIOfromRedirect(kind string, path string) IO {
	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	switch kind {
	case ">":
		return NewIO(os.Stdin, file, os.Stderr)
	case "1>":
		return NewIO(os.Stdin, file, os.Stderr)
	case "2>":
		return NewIO(os.Stdin, os.Stdout, file)
	default:
		panic("Invalid redirect kind: " + kind)
	}
}

func DefaultIO() IO {
	return NewIO(os.Stdin, os.Stdout, os.Stderr)
}

func parseArguments(input string) (string, []string, IO) {
	var name string
	var args []string
	var io IO

	if len(input) == 0 {
		panic("No command provided")
	}

	parts := split(input)

	name, parts = parts[0], parts[1:]

	if index, ok := hasRedirect(parts); ok {
		args = parts[:index]
		io = NewIOfromRedirect(parts[index], parts[index+1])
	} else {
		args = parts
		io = DefaultIO()
	}

	return name, args, io
}

func hasRedirect(parts []string) (int, bool) {
	for i, part := range parts {
		if part == "1>" || part == "2>" || part == ">" {
			return i, true
		}
	}

	return -1, false
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
