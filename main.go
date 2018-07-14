package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jessfraz/ship/aftership"
	"github.com/jessfraz/ship/version"
	"github.com/sirupsen/logrus"
)

const (
	// BANNER is what is printed for help/info output
	BANNER = `     _     _
 ___| |__ (_)_ __
/ __| '_ \| | '_ \
\__ \ | | | | |_) |
|___/_| |_|_| .__/
            |_|

 Command line tool for tracking packages using the AfterShip API.
 Version: %s
 Build: %s

`
)

var (
	apiKey string

	debug bool
	vrsn  bool
)

func init() {
	// Parse flags
	flag.StringVar(&apiKey, "apikey", os.Getenv("AFTERSHIP_API_KEY"), "AfterShip API Key (or env var AFTERSHIP_API_KEY)")

	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION, version.GITCOMMIT))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("ship version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if len(apiKey) <= 0 {
		usageAndExit("AfterShip API Key cannot be empty.", 1)
	}
}

func main() {
	// Create the AfterShip client.
	c := aftership.New(apiKey)

	// Get the trackings.
	trackings, err := c.GetTrackings()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("trackings: %#v", trackings)
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
