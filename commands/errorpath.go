package commands

import (
"fmt"
)

type ErrorPath struct {
	FilePath string
	SuppressedErrorsCount int
	OpenErrorsCount int
	ErrorsCount int
	Type string
}

func (*ErrorPath) Help() string {
	return "error path: Lists all errors found in a given file path."
}

func (*ErrorPath) Run(args []string) int {
	fmt.Printf("Error, %v", args)
	return 0
}

func (h *ErrorPath) Synopsis() string {
	return h.Help()
}