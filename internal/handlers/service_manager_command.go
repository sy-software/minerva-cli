package handlers

import (
	"fmt"

	"github.com/sy-software/minerva-cli/internal/core/ports"
	"github.com/urfave/cli/v2"
)

// ProjectManagerCommand exposes a ProjectManager instance as a CLI command
type ProjectManagerCommand struct {
	// service is a value implementing the ports.ProjectManager interface
	service ports.ProjectManager
}

// NewProjectManagerCommand initialize a ProjectManagerCommand with the given service instance
func NewProjectManagerCommand(service ports.ProjectManager) *ProjectManagerCommand {
	return &ProjectManagerCommand{
		service: service,
	}
}

// Action execute the service based on the CLI options
func (handler ProjectManagerCommand) Action(c *cli.Context) error {
	if c.Bool("init") {
		handler.service.InitRepo("NotImplemented")
	} else {
		fmt.Println("Something else")
	}

	return nil
}

// GetCommand wraps the adapter in a cli.Command struct
func (handler ProjectManagerCommand) GetCommand() *cli.Command {
	return &cli.Command{
		Name:     "project",
		Category: "backend",
		Usage:    "Creates a new backend (go lang) project",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "init",
				Usage: "Creates a new backend project",
			},
		},
		Action: handler.Action,
	}
}
