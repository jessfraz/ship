package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jessfraz/ship/aftership"
)

const shipHelp = `List shipments.`

func (cmd *listCommand) Name() string      { return "ls" }
func (cmd *listCommand) Args() string      { return "" }
func (cmd *listCommand) ShortHelp() string { return shipHelp }
func (cmd *listCommand) LongHelp() string  { return shipHelp }
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
		fmt.Printf(`%s (%s) - %s
`,
			tracking.TrackingNumber,
			tracking.Slug,
			tracking.Tag,
		)

		// Go backwards over the checkpoints so that the order is from most recent to least recent.
		for i := len(tracking.Checkpoints) - 1; i >= 0; i-- {
			location := tracking.Checkpoints[i].CountryName
			if len(tracking.Checkpoints[i].City) > 0 && len(tracking.Checkpoints[i].State) > 0 {
				location = fmt.Sprintf("%s, %s %s",
					tracking.Checkpoints[i].City,
					tracking.Checkpoints[i].State,
					tracking.Checkpoints[i].Zip,
				)
			}
			fmt.Printf(`    %s -> %s
        %s
      %s
`,
				tracking.Checkpoints[i].Tag,
				location,
				tracking.Checkpoints[i].Message,
				tracking.Checkpoints[i].CheckPointTime.Local().Format(time.RFC1123),
			)
		}
	}

	return nil
}
