package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/jessfraz/ship/aftership"
	"github.com/sirupsen/logrus"
)

const removeHelp = `Delete a shipment.`

func (cmd *removeCommand) Name() string      { return "rm" }
func (cmd *removeCommand) Args() string      { return "[OPTIONS] TRACKING_NUMBER" }
func (cmd *removeCommand) ShortHelp() string { return removeHelp }
func (cmd *removeCommand) LongHelp() string  { return removeHelp }
func (cmd *removeCommand) Hidden() bool      { return false }

func (cmd *removeCommand) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.slug, "slug", "", "slug for the carrier")
	fs.StringVar(&cmd.slug, "s", "", "slug for the carrier")
}

type removeCommand struct {
	slug string
}

func (cmd *removeCommand) Run(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass a tracking number")
	}

	cmd.slug = strings.ToLower(cmd.slug)

	if len(cmd.slug) < 1 {
		// Get the courier slug.
		courier, err := client.DetectCourier(
			aftership.Tracking{
				TrackingNumber: args[0],
			},
		)
		if err != nil {
			logrus.Fatal(err)
		}

		cmd.slug = courier.Slug
	}

	// Delete the tracking.
	if err := client.DeleteTracking(
		aftership.Tracking{
			Slug:           cmd.slug,
			TrackingNumber: args[0],
		},
	); err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("Deleted shipment for tracking number: %s\n", args[0])
	return nil
}
