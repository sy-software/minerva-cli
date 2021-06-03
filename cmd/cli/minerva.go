package main

import (
	"log"
	"os"

	"github.com/sy-software/minerva-cli/internal/core/service"
	"github.com/sy-software/minerva-cli/internal/handlers"
	"github.com/urfave/cli/v2"
)

func main() {
	service := service.NewProjectManagerService()
	handler := handlers.NewProjectManagerCommand(service)
	projectsCommand := handler.GetCommand()

	app := &cli.App{
		Name:     "Minerva CLI",
		Usage:    "A tool to manage minerva components and development workflow",
		HelpName: "minerva",
		Commands: []*cli.Command{
			projectsCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
