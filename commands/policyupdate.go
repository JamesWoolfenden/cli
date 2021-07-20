package commands

import (
"fmt"
)

type PolicyUpdate struct {

}

func (*PolicyUpdate) Help() string {
	return "Policy help"
}

func (*PolicyUpdate) Run(args []string) int {
	fmt.Printf("Policy, %v", args)
	return 0
}

func (h *PolicyUpdate) Synopsis() string {
	return h.Help()
}