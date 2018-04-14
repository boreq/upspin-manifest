package commands

import (
	"fmt"
	"github.com/boreq/guinea"
)

var buildCommit string
var buildDate string

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
	ShortDescription: "generate lists of files accesible by upspin users",
	Description:      "Upspin manifest generates lists of files accesible by upspin users.",
}

func runMain(c guinea.Context) error {
	if c.Options["version"].Bool() {
		fmt.Printf("buildCommit %s\n", buildCommit)
		fmt.Printf("buildDate %s\n", buildDate)
		return nil
	}
	return guinea.ErrInvalidParms
}
