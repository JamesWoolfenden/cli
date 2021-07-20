package commands

import (
"fmt"
)

type Scans struct {

}

func (*Scans) Help() string {
	return "Scan help"
}

func (*Scans) Run(args []string) int {
	fmt.Printf("Scans, %v", args)
	return 0
}

func (h *Scans) Synopsis() string {
	return h.Help()
}