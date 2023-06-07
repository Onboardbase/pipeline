package utils

import (
	"fmt"
	"os"

	"gopkg.in/gookit/color.v1"
)

// ErrExit Logs an error then exit the process
func ErrExit(err error, exitCode int, messages ...string) {
	if len(messages) > 0 && messages[0] != "" {
		fmt.Fprintln(os.Stderr, messages[0])
	}

	printError(err)

	if len(messages) > 0 {
		for _, message := range messages[1:] {
			fmt.Fprintln(os.Stderr, message)
		}
	}
	os.Exit(exitCode)
}

// LogError Logs an error to the console through the stdErrr
func LogError(err error) {
	printError(err)
}

func printError(e error) {
	fmt.Fprintln(os.Stderr, color.Red.Render("Pipemuta Error:"), e)
}
