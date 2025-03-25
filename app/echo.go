package main

import (
	"fmt"
	"strings"
)

func Echo(args []string, io IO) {
	fmt.Fprintln(io.Output, strings.Join(args, " "))
}
