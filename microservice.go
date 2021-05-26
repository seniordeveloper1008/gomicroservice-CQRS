package main

import (
	"bitbucket.org/burhanmubarok/microservice/instances/infrastructures/commands"
	"bitbucket.org/burhanmubarok/microservice/interfaces/infrastructures/commands"
)

func main() {
	var cli interfaces.CLI
	cli = &instances.CLI{}
	cli.Execute()
}
