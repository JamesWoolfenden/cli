package commands

import (
"fmt"
)

type PolicyAdd struct {

}

func (*PolicyAdd) Help() string {
	return "Adds a Policy to the Bridgecrew Platform"
}

func (*PolicyAdd) Run(args []string) int {
	fmt.Printf("Policy, %v", args)
	return 0
}

func (h *PolicyAdd) Synopsis() string {
	return h.Help()
}