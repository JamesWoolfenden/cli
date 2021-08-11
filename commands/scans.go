package commands

import (
	"fmt"
)

const url = "https://www.bridgecrew.cloud/api/v1/scans/integrations"

type Scans struct {
}

func (*Scans) Help() string {
	return "Invokes a bridgecrew platform repository scan"
}

func (*Scans) Run(args []string) int {
	fmt.Printf("Invoked")
	return 0
}

func (h *Scans) Synopsis() string {
	return h.Help()
}
