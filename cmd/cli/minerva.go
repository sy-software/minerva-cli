package main

import (
	"log"
	"os"

	"github.com/sy-software/minerva-cli/internal/core/service"
	"github.com/sy-software/minerva-cli/internal/handlers"
	"github.com/urfave/cli/v2"
)

// Exposes minerva as a CLI
func main() {
	projService := service.NewProjectManagerService()
	projHandler := handlers.NewProjectManagerCommand(projService)
	projCommand := projHandler.GetCommand()

	testService := &service.TestRunnerService{}
	testHandler := handlers.NewTestRunnerCommand(testService)
	testCommand := testHandler.GetCommand()

	app := &cli.App{
		Name:     "Minerva CLI",
		Usage:    "A tool to manage minerva components and development workflow",
		HelpName: "minerva",
		Commands: []*cli.Command{
			projCommand,
			testCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
