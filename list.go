package main

import (
	"flag"
	"fmt"

	"github.com/jessfraz/ship/aftership"
	"github.com/sirupsen/logrus"
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

	// Go backwards over the trackings so that the order is from most recent to least recent.
	for i := len(trackings) - 1; i >= 0; i-- {
		prettyPrintTracking(trackings[i], false)
		fmt.Println()
	}

	return nil
}
