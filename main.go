package main

import (
	"fmt"
	"os"

	"github.com/jameswoolfenden/bridge/commands"
	"github.com/mitchellh/cli"
)

func main() {
	app := cli.NewCLI("bridge", "0.0.0")
	app.Args = os.Args[1:]
	app.Commands = map[string]cli.CommandFactory{
		"repo list": func() (cli.Command, error) {
			return &commands.Repo{}, nil
		},
		"repo update": func() (cli.Command, error) {
			return &commands.Repo{}, nil
		},
		"error list": func() (cli.Command, error) {
			return &commands.Error{}, nil
		},
		"error list authors": func() (cli.Command, error) {
			return &commands.Error{}, nil
		},
		"policy validate": func() (cli.Command, error) {
			return &commands.Policy{}, nil
		},
		"policy add": func() (cli.Command, error) {
			return &commands.Policy{}, nil
		},
		"policy delete": func() (cli.Command, error) {
			return &commands.Policy{}, nil
		},
		"policy update": func() (cli.Command, error) {
			return &commands.Policy{}, nil
		},
		"scan trigger": func() (cli.Command, error) {
			return &commands.Scans{}, nil
		},
		"suppression trigger": func() (cli.Command, error) {
			return &commands.Suppression{}, nil
		},
		"fix add": func() (cli.Command, error) {
			return &commands.Suppression{}, nil
		},
	}

	status, err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}

type Bridge struct {
}

func (*Bridge) Help() string {
	return "bridge help"
}

func (*Bridge) Run(args []string) int {
	fmt.Printf("bridge, %v", args)
	return 0
}

func (h *Bridge) Synopsis() string {
	return h.Help()
}
