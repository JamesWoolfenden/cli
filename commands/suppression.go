package commands

import (
"fmt"
)

type Suppression struct {

}

func (*Suppression) Help() string {
	return "Suppression help"
}

func (*Suppression) Run(args []string) int {
	fmt.Printf("Suppression, %v", args)
	return 0
}

func (h *Suppression) Synopsis() string {
	return h.Help()
}