// Package manifest handles manifest file decoding.
package manifest

import (
	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Path      string                    `yaml:"path"`
	Manifests map[string]ManifestConfig `yaml:"manifests"`
}

type ManifestConfig struct {
	Header *string  `yaml:"header"`
	Users  []string `yaml:"users"`
}

func Load(data []byte) (Manifest, error) {
	var manifest Manifest

	err := yaml.Unmarshal(data, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}
