// Package builder produces and puts the manifest files into upspin.
package builder

import (
	"bytes"
	"fmt"
	"github.com/boreq/upspin-manifest/manifest"
	"github.com/boreq/upspin-manifest/upspin"
	"sort"
)

type Builder interface {
	Build(userFiles map[string][]string, userDirectories map[string][]string, man manifest.Manifest) error
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

// Build produces the manifest files and puts them into upspin as specified by
// the manifest file. It accepts a map mapping usernames to accessible
// filenames as the first argument. This will most likely be produced by the
// parser.
func (b *builder) Build(userFiles map[string][]string, userDirectories map[string][]string, man manifest.Manifest) error {
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
			out.Write([]byte(fmt.Sprintf("%s:", user)))

			// Directories
			if manConfig.ListDirectories == nil || *manConfig.ListDirectories {
				directories := userDirectories[user]
				if len(directories) > 0 {
					out.Write([]byte("\n"))
				}

				sort.Slice(directories, func(i, j int) bool { return directories[i] < directories[j] })

				for _, directory := range directories {
					out.Write([]byte("\t"))
					out.Write([]byte(directory))
					out.Write([]byte("\n"))
				}
			}

			// Files
			if manConfig.ListFiles == nil || *manConfig.ListFiles {
				files := userFiles[user]
				if len(files) > 0 {
					out.Write([]byte("\n"))
				}

				sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })

				for _, file := range files {
					out.Write([]byte("\t"))
					out.Write([]byte(file))
					out.Write([]byte("\n"))
				}
			}
		}

		err := b.ups.Put(target, out.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}
