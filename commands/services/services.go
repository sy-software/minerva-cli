// Package services CLI sub-commands to handle actions related to services
package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
	"sysoft.com/minerva-cli/utils"
)

// Base template address for a service repository
const SERVICE_TEMPLATE = "git@github.com:sy-software/minerva-service-template.git"

// GitNotFound indicates git binary is not installed in the machine
var GitNotFound = errors.New("git command not found")

// cloneTemplate clones the service template repository
// returns an error if git not installed in the machine
func cloneTemplate() (string, error) {
	tempDir := fmt.Sprint(os.TempDir(), "service-template-", time.Now().UnixNano())
	gitPath, err := exec.LookPath("git")

	if err != nil {
		return tempDir, GitNotFound
	}

	clone := &exec.Cmd{
		Path:   gitPath,
		Args:   []string{gitPath, "clone", SERVICE_TEMPLATE, tempDir},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	err = clone.Run()
	return tempDir, err
}

/*
ServiceCommand execute all available actions for minerva services

Action Init

Init action clones the service template into the current directory
*/
var ServiceCommand = &cli.Command{
	Name:     "service",
	Category: "template",
	Usage:    "Manages service projects",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "init",
			Usage: "Creates a new service project",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("init") {
			utils.Info("Initializing new service project")
			templatePath, err := cloneTemplate()

			if err != nil {
				log.Fatal(err)
			}

			utils.CopyDir(templatePath, "output/")
		} else {
			fmt.Println("Something else")
		}

		return nil
	},
}
