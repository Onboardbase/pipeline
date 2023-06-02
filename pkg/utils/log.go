package utils

import (
	"fmt"
	"os"

	"gopkg.in/gookit/color.v1"
)

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
func LogError(err error) {
	printError(err)
}

func printError(e error) {
	fmt.Fprintln(os.Stderr, color.Red.Render("CLI Tools Error:"), e)
}
