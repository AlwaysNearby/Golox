package golox

import "fmt"

var HadError bool

func sendError(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Printf("[line %d] Error%s: %s\n", line, where, message)
	HadError = true
}
