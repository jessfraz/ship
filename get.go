package main

import (
	"context"
	"errors"
	"flag"
	"strings"

	"github.com/jessfraz/ship/aftership"
	"github.com/sirupsen/logrus"
)

const getHelp = `Get details for a shipment.`

func (cmd *getCommand) Name() string      { return "get" }
func (cmd *getCommand) Args() string      { return "[OPTIONS] TRACKING_NUMBER" }
func (cmd *getCommand) ShortHelp() string { return getHelp }
func (cmd *getCommand) LongHelp() string  { return getHelp }
func (cmd *getCommand) Hidden() bool      { return false }

func (cmd *getCommand) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.slug, "slug", "", "slug for the carrier")
	fs.StringVar(&cmd.slug, "s", "", "slug for the carrier")
}

type getCommand struct {
	slug string
}

func (cmd *getCommand) Run(ctx context.Context, args []string) error {
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

	// Get the tracking.
	tracking, err := client.GetTracking(
		aftership.Tracking{
			Slug:           cmd.slug,
			TrackingNumber: args[0],
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	prettyPrintTracking(tracking, true)

	return nil
}
