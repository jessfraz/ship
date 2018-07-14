package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/jessfraz/ship/aftership"
)

const listHelp = `List shipments.`

func (cmd *listCommand) Name() string      { return "ls" }
func (cmd *listCommand) Args() string      { return "" }
func (cmd *listCommand) ShortHelp() string { return listHelp }
func (cmd *listCommand) LongHelp() string  { return listHelp }
func (cmd *listCommand) Hidden() bool      { return false }

func (cmd *listCommand) Register(fs *flag.FlagSet) {}

type listCommand struct{}

func (cmd *listCommand) Run(c *aftership.Client, args []string) error {
	// Get the trackings.
	trackings, err := c.GetTrackings()
	if err != nil {
		logrus.Fatal(err)
	}

	for _, tracking := range trackings {
		prettyPrintTracking(tracking)
	}

	return nil
}
