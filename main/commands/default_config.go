package commands

import (
	"encoding/json"
	"fmt"
	"github.com/boreq/guinea"
	"github.com/boreq/upspin-manifest/config"
)

var defaultConfigCmd = guinea.Command{
	Run:              runDefaultConfig,
	ShortDescription: "prints the default configuration to stdout",
	Description: `
This command prints out the default config in the json format to stdout. The
available config keys are:

Debug
	Specifies if the program should run in the debug mode. The program
	running in the debug mode prints more log messages.
	Allowed values: true or false.`,
}

func runDefaultConfig(c guinea.Context) error {
	defaultConfig := config.Default()
	j, err := json.MarshalIndent(defaultConfig, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(j))
	return nil
}
