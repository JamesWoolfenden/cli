package commands

import (
	"fmt"
)

type Policy struct {
}

func (*Policy) Help() string {
	return "Validates a Policy"
}

func (*Policy) Run(args []string) int {
	fmt.Printf("Policy, %v", args)
	return 0
}

func (h *Policy) Synopsis() string {
	return h.Help()
}
