package commands

import (
"fmt"
)

type RepoBranches struct {

}

func (*RepoBranches) Help() string {
	return "Get the repository branches (that has CICD runs)"
}

func (*RepoBranches) Run(args []string) int {
	fmt.Printf("repo, %v", args)
	return 0
}

func (h *RepoBranches) Synopsis() string {
	return h.Help()
}