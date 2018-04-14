package builder

import (
	"bytes"
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
	for target, manConfig := range man.Manifests {
		var out bytes.Buffer
		if manConfig.Header != nil {
			header, err := b.ups.Get(*manConfig.Header)
			if err != nil {
				return err
			}
			out.Write(header)
			out.Write([]byte("\n"))
		}

		for _, user := range manConfig.Users {
			for _, file := range userFiles[user] {
				out.Write([]byte(file))
				out.Write([]byte("\n"))
			}
		}

		err := b.ups.Put(target, out.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}
