package main

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/genuinetools/pkg/cli"
	"github.com/jessfraz/ship/aftership"
	"github.com/jessfraz/ship/version"
	"github.com/sirupsen/logrus"
)

var (
	apiKey string
	debug  bool

	client *aftership.Client
)

func main() {
	// Create a new cli program.
	p := cli.NewProgram()
	p.Name = "ship"
	p.Description = "Command line tool to track packages using the AfterShip API"
	// Set the GitCommit and Version.
	p.GitCommit = version.GITCOMMIT
	p.Version = version.VERSION

	// Build the list of available commands.
	p.Commands = []cli.Command{
		&createCommand{},
		&getCommand{},
		&listCommand{},
		&removeCommand{},
	}

	// Setup the global flags.
	p.FlagSet = flag.NewFlagSet("ship", flag.ExitOnError)
	p.FlagSet.BoolVar(&debug, "d", false, "enable debug logging")
	p.FlagSet.StringVar(&apiKey, "apikey", os.Getenv("AFTERSHIP_API_KEY"), "AfterShip API Key (or env var AFTERSHIP_API_KEY)")

	// Set the before function.
	p.Before = func(ctx context.Context) error {
		// Set the log level.
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		// Only do the next thing if we actually have arguments.
		if len(os.Args) > 1 {
			if len(apiKey) <= 0 && os.Args[1] != "version" {
				return errors.New("AfterShip API Key cannot be empty")
			}

			// Create the AfterShip client.
			client = aftership.New(apiKey)
		}

		return nil
	}

	// Run our program.
	p.Run()
}
