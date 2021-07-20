package commands

import (
"fmt"
)

type Fixes struct {

}

func (*Fixes) Help() string {
	return "Fixes help"
}

func (*Fixes) Run(args []string) int {
	fmt.Printf("Fixes, %v", args)
	return 0
}

func (h *Fixes) Synopsis() string {
	return h.Help()
}