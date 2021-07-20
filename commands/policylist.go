package commands

import (
"fmt"
)

type PolicyList struct {

}

func (*PolicyList) Help() string {
	return "Lists all Custom Policies"
}

func (*PolicyList) Run(args []string) int {
	fmt.Printf("Policy, %v", args)
	return 0
}

func (h *PolicyList) Synopsis() string {
	return h.Help()
}