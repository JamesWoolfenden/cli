package commands

import (
	"fmt"
)

type PolicyDelete struct {
}

func (*PolicyDelete) Help() string {
	return "Delete a Policy"
}

func (*PolicyDelete) Run(args []string) int {
	fmt.Printf("Policy, %v", args)
	return 0
}

func (h *PolicyDelete) Synopsis() string {
	return h.Help()
}
