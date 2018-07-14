package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/jessfraz/ship/aftership"
	"github.com/jessfraz/ship/version"
)

const versionHelp = `Show the version information.`

func (cmd *versionCommand) Name() string       { return "version" }
func (cmd *versionCommand) Args() string       { return "" }
func (cmd *versionCommand) ShortHelp() string  { return versionHelp }
func (cmd *versionCommand) LongHelp() string   { return versionHelp }
func (cmd *versionCommand) Hidden() bool       { return false }
func (cmd *versionCommand) DoReexec() bool     { return false }
func (cmd *versionCommand) RequiresRunc() bool { return false }

func (cmd *versionCommand) Register(fs *flag.FlagSet) {}

type versionCommand struct{}

func (cmd *versionCommand) Run(c *aftership.Client, args []string) error {
	fmt.Printf(`%s:
 version     : %s
 git hash    : %s
 go version  : %s
 go compiler : %s
 platform    : %s/%s
`, "ship", version.VERSION, version.GITCOMMIT,
		runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
	return nil
}
