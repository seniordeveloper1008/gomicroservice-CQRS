package main

import (
	instances "github.com/burhanmubarok/microservice/instances/infrastructures/commands"
	interfaces "github.com/burhanmubarok/microservice/interfaces/infrastructures/commands"
)

func main() {
	var cli interfaces.CLI
	cli = &instances.CLI{}
	cli.Execute()
}
