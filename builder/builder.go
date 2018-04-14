package builder

import (
	"fmt"
	"github.com/boreq/upspin-manifest/manifest"
	"github.com/boreq/upspin-manifest/upspin"
)

type Builder interface {
	Build(userFiles map[string][]string, man manifest.Manifest) error
}

func New(ups upspin.Upspin) Builder {
	rv := &builder{
		ups: ups,
	}
	return rv
}

type builder struct {
	ups upspin.Upspin
}

func (b *builder) Build(userFiles map[string][]string, man manifest.Manifest) error {
	for user, files := range userFiles {
		for _, f := range files {
			fmt.Printf("%s %s\n", user, f)
		}
	}

	return nil
}
