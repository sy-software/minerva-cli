package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"sysoft.com/minerva-cli/commands/services"
)

func main() {
	app := &cli.App{
		Name:  "Minerva CLI",
		Usage: "A tool to manage minerva components and development workflow",
		Commands: []*cli.Command{
			services.ServiceCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
