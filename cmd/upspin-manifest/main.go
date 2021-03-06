// Package main contains the main function for upspin-manifest.
package main

import (
	"fmt"
	"github.com/boreq/guinea"
	"github.com/boreq/upspin-manifest/cmd/upspin-manifest/commands"
	"os"
)

func main() {
	e := guinea.Run(&commands.MainCmd)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
