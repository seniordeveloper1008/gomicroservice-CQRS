package instances

import (
	interfaces "github.com/burubur/microservice/interfaces/infrastructures/commands"
	"github.com/spf13/cobra"
)

// CLI doc
type CLI struct{}

// NewCLI doc
func NewCLI() *CLI {
	return &CLI{}
}

// Execute doc
func (c *CLI) Execute() {
	cli := &cobra.Command{
		Use:   "microservice",
		Short: "Microservice <name>",
		Long:  "Microservice <name>. This microservice will be running under CLI, look at the Usage and Available Commands below for the detail.",
	}

	var microservice interfaces.Microservice
	microservice = &Microservice{}

	cliCommands := []*cobra.Command{
		{
			Use:   "ensure",
			Short: "Configuration Check",
			Long:  "a task that will be ensuring configuration file and it's value.",
			Run: func(_ *cobra.Command, _ []string) {
				microservice.Ensure()
			},
		},
		{
			Use:   "test",
			Short: "Connection Test",
			Long:  "a task that will assess that all required configuration has been complied and test some of it's module (database, email, kafka, log, etc).",
			Run: func(_ *cobra.Command, _ []string) {
				microservice.Test()
			},
		},
		{
			Use:   "migrate",
			Short: "Database Migration",
			Long:  "a task that will be migrating and seeding database.",
			Run: func(_ *cobra.Command, _ []string) {
				microservice.Migrate()
			},
		},
		{
			Use:   "serve",
			Short: "Server Listener",
			Long:  "a server listener that will be listening any http request.",
			Run: func(_ *cobra.Command, _ []string) {
				microservice.Serve()
			},
		},
	}

	for _, command := range cliCommands {
		cli.AddCommand(command)
	}
	cli.Execute()
}
