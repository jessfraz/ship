package main

import (
	"errors"
	"flag"

	"github.com/jessfraz/ship/aftership"
	"github.com/sirupsen/logrus"
)

const getHelp = `Get details for a shipment.`

func (cmd *getCommand) Name() string      { return "get" }
func (cmd *getCommand) Args() string      { return "[OPTIONS] TRACKING_NUMBER" }
func (cmd *getCommand) ShortHelp() string { return getHelp }
func (cmd *getCommand) LongHelp() string  { return getHelp }
func (cmd *getCommand) Hidden() bool      { return false }

func (cmd *getCommand) Register(fs *flag.FlagSet) {}

type getCommand struct{}

func (cmd *getCommand) Run(c *aftership.Client, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass a tracking number")
	}

	// Get the courier slug.
	courier, err := c.DetectCourier(
		aftership.Tracking{
			TrackingNumber: args[0],
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	// Get the tracking.
	tracking, err := c.GetTracking(
		aftership.Tracking{
			Slug:           courier.Slug,
			TrackingNumber: args[0],
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	prettyPrintTracking(tracking, true)

	return nil
}
