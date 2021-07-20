package commands

import (
"fmt"
)

type PolicyPreview struct {

}

func (*PolicyPreview) Help() string {
	return "Policy Preview"
}

func (*PolicyPreview) Run(args []string) int {
	fmt.Printf("Policy, %v", args)
	return 0
}

func (h *PolicyPreview) Synopsis() string {
	return h.Help()
}