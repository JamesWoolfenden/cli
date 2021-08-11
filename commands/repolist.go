package commands

import (
	"fmt"
)

type Repo struct {
}

func (*Repo) Help() string {
	return "Get a complete repository list"
}

func (*Repo) Run(args []string) int {
	fmt.Printf("repo, %v", args)
	return 0
}

func (h *Repo) Synopsis() string {
	return h.Help()
}
