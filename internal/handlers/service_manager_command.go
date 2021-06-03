package handlers

import (
	"fmt"

	"github.com/sy-software/minerva-cli/internal/core/ports"
	"github.com/urfave/cli/v2"
)

type ProjectManagerCommand struct {
	service ports.ProjectManager
}

func NewProjectManagerCommand(service ports.ProjectManager) *ProjectManagerCommand {
	return &ProjectManagerCommand{
		service: service,
	}
}

func (handler ProjectManagerCommand) Action(c *cli.Context) error {
	if c.Bool("init") {
		handler.service.InitRepo("NotImplemented")
	} else {
		fmt.Println("Something else")
	}

	return nil
}

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
