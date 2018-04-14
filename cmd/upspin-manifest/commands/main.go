package commands

import (
	"fmt"
	"github.com/boreq/guinea"
	"github.com/boreq/upspin-manifest/cmd"
)

var MainCmd = guinea.Command{
	Options: []guinea.Option{
		guinea.Option{
			Name:        "version",
			Type:        guinea.Bool,
			Description: "Display version",
		},
	},
	Run: runMain,
	Subcommands: map[string]*guinea.Command{
		"run":            &runCmd,
		"default_config": &defaultConfigCmd,
	},
	ShortDescription: "SDR plane tracking software",
	Description:      "This software records plane tracking data collected by SDR radios.",
}

func runMain(c guinea.Context) error {
	if c.Options["version"].Bool() {
		fmt.Printf("BuildCommit %s\n", cmd.BuildCommit)
		fmt.Printf("BuildDate %s\n", cmd.BuildDate)
		return nil
	}
	return guinea.ErrInvalidParms
}
