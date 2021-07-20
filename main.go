package main

import (
	"fmt"
	"os"
	"log"

	"github.com/jameswoolfenden/bridge/commands"
	"github.com/mitchellh/cli"
)

func main() {
	app := cli.NewCLI("bridge", "0.0.1")

	api := os.Getenv("BRIDGECREW_API")
    
    if api == "" {
        log.Fatal("BRIDGECREW_API is missing")
    }

	app.Args = os.Args[1:]
	app.Commands = map[string]cli.CommandFactory{
		"repo list": func() (cli.Command, error) {
			return &commands.Repo{}, nil
		},
		"repo branches": func() (cli.Command, error) {
			return &commands.RepoBranches{}, nil
		},
		"error list": func() (cli.Command, error) {
			return &commands.ErrorList{}, nil
		},
		"error path": func() (cli.Command, error) {
			return &commands.ErrorPath{}, nil
		},
		"error list authors": func() (cli.Command, error) {
			return &commands.ErrorList{}, nil
		},
		"policy validate": func() (cli.Command, error) {
			return &commands.Policy{}, nil
		},
		"policy list": func() (cli.Command, error) {
			return &commands.PolicyList{}, nil
		},
		"policy add": func() (cli.Command, error) {
			return &commands.PolicyAdd{}, nil
		},
		"policy delete": func() (cli.Command, error) {
			return &commands.PolicyDelete{}, nil
		},
		"policy update": func() (cli.Command, error) {
			return &commands.PolicyUpdate{}, nil
		},
		"policy preview": func() (cli.Command, error) {
			return &commands.PolicyPreview{}, nil
		},
		"scan trigger": func() (cli.Command, error) {
			return &commands.Scans{}, nil
		},
		"suppression list": func() (cli.Command, error) {
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
