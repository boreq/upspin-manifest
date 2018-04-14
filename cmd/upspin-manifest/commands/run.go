package commands

import (
	"github.com/boreq/guinea"
	"github.com/boreq/upspin-manifest/builder"
	"github.com/boreq/upspin-manifest/config"
	"github.com/boreq/upspin-manifest/logging"
	"github.com/boreq/upspin-manifest/manifest"
	"github.com/boreq/upspin-manifest/parser"
	"github.com/boreq/upspin-manifest/upspin"
)

var log = logging.GetLogger("run")

var runCmd = guinea.Command{
	Run: runRun,
	Options: []guinea.Option{
		guinea.Option{
			Name:        "config",
			Type:        guinea.String,
			Description: "Config file",
		},
	},
	Arguments: []guinea.Argument{
		guinea.Argument{
			Name:        "manifest",
			Multiple:    false,
			Description: "Manifest file",
		},
	},
	ShortDescription: "runs the program",
}

func runRun(c guinea.Context) error {
	// Load the config if requested
	if c.Options["config"].Str() != "" {
		if err := config.Load(c.Options["config"].Str()); err != nil {
			return err
		}
	}

	ups := upspin.New(config.Config.UpspinExecutable)

	// Load the manifest
	log.Debug("Loading the manifest...")
	manData, err := ups.Get(c.Arguments[0])
	if err != nil {
		return err
	}
	man, err := manifest.Load(manData)
	if err != nil {
		return err
	}

	// Load the share data
	log.Debug("Loading the share data...")
	shareData, err := ups.Share(man.Path)
	if err != nil {
		return err
	}
	userFiles, userDirectories, err := parser.Parse(shareData)
	if err != nil {
		return err
	}

	// Build and write the file lists
	log.Debug("Building...")
	buil := builder.New(ups)
	err = buil.Build(userFiles, userDirectories, man)
	if err != nil {
		return err
	}

	return nil
}
