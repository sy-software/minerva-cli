package handlers

import (
	"github.com/sy-software/minerva-cli/internal/core/ports"
	"github.com/urfave/cli/v2"
)

type TestRunnerCommand struct {
	service ports.TestRunner
}

func NewTestRunnerCommand(service ports.TestRunner) *TestRunnerCommand {
	return &TestRunnerCommand{
		service: service,
	}
}

// Action execute the service based on the CLI options
func (handler TestRunnerCommand) Action(c *cli.Context) error {
	path := c.Args().Get(0)

	if c.Bool("discover") {
		handler.service.DiscoverTests(path)
	} else {
		watch := c.Bool("watch")
		// args []string
		args := []string{}
		if c.Args().Len() > 1 {
			args = c.Args().Slice()[1:]
		}

		err := handler.service.RunTests(path, watch, args)

		if err != nil {
			// Exit code will be printed with the same message so we can return
			// an empty value
			return cli.Exit("", 1)
		}
	}

	return nil
}

// GetCommand wraps the adapter in a cli.Command struct
func (handler TestRunnerCommand) GetCommand() *cli.Command {
	return &cli.Command{
		Name:        "test",
		Category:    "util",
		Usage:       "An utility to run go native test",
		ArgsUsage:   "<path>",
		Description: "<path>: The path to look for tests",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "watch",
				Aliases: []string{"w"},
				Value:   false,
				Usage:   "Run tests in watch mode",
			},
			&cli.BoolFlag{
				Name:    "discover",
				Aliases: []string{"d"},
				Value:   false,
				Usage:   "Discovers and list existing tests",
			},
		},
		Action: handler.Action,
	}
}
