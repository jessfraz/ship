package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jessfraz/ship/aftership"
)

const createHelp = `Create a shipment.`

func (cmd *createCommand) Name() string      { return "create" }
func (cmd *createCommand) Args() string      { return "[OPTIONS] TRACKING_NUMBER" }
func (cmd *createCommand) ShortHelp() string { return createHelp }
func (cmd *createCommand) LongHelp() string  { return createHelp }
func (cmd *createCommand) Hidden() bool      { return false }

func (cmd *createCommand) Register(fs *flag.FlagSet) {}

type createCommand struct{}

func (cmd *createCommand) Run(c *aftership.Client, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass a tracking number")
	}

	// Create the tracking.
	tracking, err := c.PostTracking(
		aftership.Tracking{
			TrackingNumber: args[0],
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	prettyPrintTracking(tracking)

	return nil
}

func prettyPrintTracking(tracking aftership.Tracking) {
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
