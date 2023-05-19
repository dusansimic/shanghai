package libshg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Shanghaifile struct {
	Tree   Node        `yaml:"tree"`
	Images MapOfImages `yaml:"images"`
}

type Node map[string]interface{}

type MapOfImages map[string]Image

type Image struct {
	Tag           string   `yaml:"tag"`
	ContainerFile string   `yaml:"containerfile"`
	Context       string   `yaml:"context"`
	BuildArgs     []string `yaml:"buildargs"`
}

func ReadShanghaifile(f string) (*Shanghaifile, error) {
	d, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read Shanghaifile: %w", err)
	}

	shg := &Shanghaifile{}
	if yaml.Unmarshal(d, &shg) != nil {
		return nil, fmt.Errorf("failed to parse Shanghaifile: %w", err)
	}

	return shg, nil
}
