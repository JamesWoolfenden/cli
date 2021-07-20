package commands

import (
"fmt"
)

type Repo struct {

}

func (*Repo) Help() string {
	return "repo help"
}

func (*Repo) Run(args []string) int {
	fmt.Printf("repo, %v", args)
	return 0
}

func (h *Repo) Synopsis() string {
	return h.Help()
}