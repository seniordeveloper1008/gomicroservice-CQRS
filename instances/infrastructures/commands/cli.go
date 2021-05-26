package instances

import (
	"bitbucket.org/burhanmubarok/microservice/interfaces/infrastructures/commands"
	"github.com/spf13/cobra"
)

// CLI doc
type CLI struct{}

// Execute doc
func (c *CLI) Execute() {
	cli := &cobra.Command{
		Use:   "Microservice",
		Short: "Microservice [name]",
		Long:  "Microservice [desc]",
	}

	var microservice interfaces.Microservice
	microservice = &Microservice{}

	cliCommands := []*cobra.Command{
		{
			Use:   "ensure",
			Short: "Configuration Check",
			Long:  "a server listener that ran by command line interface",
			Run: func(command *cobra.Command, args []string) {
				microservice.Ensure()
			},
		},
		{
			Use:   "test",
			Short: "Connection Test",
			Long:  "a tester to assure that this microservice have a connection and a certain access (database, email, kafka, log, etc)",
			Run: func(command *cobra.Command, args []string) {
				microservice.Test()
			},
		},
		{
			Use:   "migrate",
			Short: "Database Migration",
			Long:  "a server listener that ran by command line interface",
			Run: func(command *cobra.Command, args []string) {
				microservice.Migrate()
			},
		},
		{
			Use:   "serve",
			Short: "Server Listener",
			Long:  "a server listener that ran by command line interface",
			Run: func(command *cobra.Command, args []string) {
				microservice.Serve()
			},
		},
	}

	for _, command := range cliCommands {
		cli.AddCommand(command)
	}
	cli.Execute()
}
