package commands

import (
"fmt"
)

type Error struct {

}

func (*Error) Help() string {
	return "error help"
}

func (*Error) Run(args []string) int {
	fmt.Printf("Error, %v", args)
	return 0
}

func (h *Error) Synopsis() string {
	return h.Help()
}