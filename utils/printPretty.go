package utils

import (
	"fmt"
	"strings"
)

// PrintPretty print a string with a pretty format
func PrintPretty(msg string) string {
	line := strings.Repeat("*", 2*len(msg))
	out := fmt.Sprintf("\n")
	out += fmt.Sprintf("%s\n", line)
	out += fmt.Sprintf("\t\t%s\n", msg)
	out += fmt.Sprintf("%s\n", line)
	return out
}
