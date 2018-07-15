package main

import (
	"context"
	"flag"
	"fmt"

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

func (cmd *listCommand) Run(ctx context.Context, args []string) error {
	// Get the trackings.
	trackings, err := client.GetTrackings()
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
