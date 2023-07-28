package shanghai

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
	Tag           string            `yaml:"tag"`
	ContainerFile string            `yaml:"containerfile"`
	Context       string            `yaml:"context"`
	BuildArgs     map[string]string `yaml:"buildargs"`
}

type shanghaifile struct {
	Tree   Node `yaml:"tree"`
	Images map[string]struct {
		Tag           string                 `yaml:"tag"`
		ContainerFile string                 `yaml:"containerfile"`
		Context       string                 `yaml:"context"`
		BuildArgs     map[string]interface{} `yaml:"buildargs"`
	} `yaml:"images"`
}

func ReadShanghaifile(f string) (*Shanghaifile, error) {
	d, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read Shanghaifile: %w", err)
	}

	siface := &shanghaifile{}
	if yaml.Unmarshal(d, &siface) != nil {
		return nil, fmt.Errorf("failed to unmarshal Shanghaifile: %w", err)
	}

	s := &Shanghaifile{
		Tree:   siface.Tree,
		Images: make(MapOfImages),
	}

	for k, v := range siface.Images {
		s.Images[k] = Image{
			Tag:           v.Tag,
			ContainerFile: v.ContainerFile,
			Context:       v.Context,
			BuildArgs:     make(map[string]string),
		}

		for k2, v2 := range v.BuildArgs {
			switch sv := v2.(type) {
			case string:
				s.Images[k].BuildArgs[k2] = sv
			default:
				return nil, fmt.Errorf("incorret value in buildargs (key: %s): %w", k2, err)
			}
		}
	}

	return s, nil
}
