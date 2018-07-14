package main

import (
	"errors"
	"flag"

	"github.com/jessfraz/ship/aftership"
	"github.com/sirupsen/logrus"
)

const removeHelp = `Delete a shipment.`

func (cmd *removeCommand) Name() string      { return "rm" }
func (cmd *removeCommand) Args() string      { return "[OPTIONS] TRACKING_NUMBER" }
func (cmd *removeCommand) ShortHelp() string { return removeHelp }
func (cmd *removeCommand) LongHelp() string  { return removeHelp }
func (cmd *removeCommand) Hidden() bool      { return false }

func (cmd *removeCommand) Register(fs *flag.FlagSet) {}

type removeCommand struct{}

func (cmd *removeCommand) Run(c *aftership.Client, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass a tracking number")
	}

	// remove the courier slug.
	courier, err := c.DetectCourier(
		aftership.Tracking{
			TrackingNumber: args[0],
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	// Delete the tracking.
	if err := c.DeleteTracking(
		aftership.Tracking{
			Slug:           courier.Slug,
			TrackingNumber: args[0],
		},
	); err != nil {
		logrus.Fatal(err)
	}

	return nil
}
