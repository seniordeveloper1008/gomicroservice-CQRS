package main

import (
	instances "github.com/burubur/microservice/instances/infrastructures/commands"
	interfaces "github.com/burubur/microservice/interfaces/infrastructures/commands"
)

func main() {
	var cli interfaces.CLI
	cli = instances.NewCLI()
	cli.Execute()
}
