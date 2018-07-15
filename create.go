package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jessfraz/ship/aftership"
	"github.com/sirupsen/logrus"
)

const createHelp = `Create a shipment.`

func (cmd *createCommand) Name() string      { return "create" }
func (cmd *createCommand) Args() string      { return "[OPTIONS] TRACKING_NUMBER" }
func (cmd *createCommand) ShortHelp() string { return createHelp }
func (cmd *createCommand) LongHelp() string  { return createHelp }
func (cmd *createCommand) Hidden() bool      { return false }

func (cmd *createCommand) Register(fs *flag.FlagSet) {}

type createCommand struct{}

func (cmd *createCommand) Run(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass a tracking number")
	}

	// Create the tracking.
	tracking, err := client.PostTracking(
		aftership.Tracking{
			TrackingNumber: args[0],
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	if len(tracking.TrackingNumber) <= 0 {
		fmt.Printf("Created shipment for tracking number: %s\n", args[0])
		return nil
	}

	prettyPrintTracking(tracking, true)

	return nil
}

func prettyPrintTracking(tracking aftership.Tracking, all bool) {
	fmt.Printf(`%s (%s) - %s
`,
		tracking.TrackingNumber,
		color.HiBlackString(strings.ToUpper(tracking.Slug)),
		printTag(tracking.Tag),
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
			printTag(tracking.Checkpoints[i].Tag),
			location,
			tracking.Checkpoints[i].Message,
			color.BlackString(tracking.Checkpoints[i].CheckPointTime.Local().Format(time.RFC1123)),
		)
		// If we don't want to print all the checkpoints set i=0 so we don't..
		if !all {
			i = 0
		}
	}
}

// Statuses come from: https://docs.aftership.com/api/4/delivery-status
func printTag(tag string) string {
	switch tag {
	case "InTransit":
		return color.YellowString(tag)
	case "Delivered":
		return color.GreenString(tag)
	case "OutForDelivery":
		return color.CyanString(tag)
	case "InfoReceived":
		return color.BlueString(tag)
	case "Exception":
		return color.MagentaString(tag)
	case "FailedAttempt":
		return color.RedString(tag)
	case "Expired":
		return color.MagentaString(tag)
	case "Pending":
		return color.HiWhiteString(tag)
	}
	return color.WhiteString(tag)
}
