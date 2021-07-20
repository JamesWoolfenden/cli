package commands

import (
"fmt"
)

type ErrorList struct {
	FilePath string
	SuppressedErrorsCount int
	OpenErrorsCount int
	ErrorsCount int
	Type string
}

func (*ErrorList) Help() string {
	return "Lists scanned files that contain errors."
}

func (*ErrorList) Run(args []string) int {
	fmt.Printf("Error, %v", args)
	return 0
}

func (h *ErrorList) Synopsis() string {
	return h.Help()
}