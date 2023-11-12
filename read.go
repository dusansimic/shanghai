package shanghai

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Shanghaifile struct {
	Tree                 Tree
	EnvironmentVariables map[string]string
	BuildArguments       map[string]string
}

type treeNode map[string]interface{}

type MapOfImages map[string]Image

type shanghaifile struct {
	Tree   treeNode `yaml:"tree"`
	Images map[string]struct {
		Tag           string                 `yaml:"tag"`
		ContainerFile string                 `yaml:"containerfile"`
		Context       string                 `yaml:"context"`
		BuildArgs     map[string]interface{} `yaml:"buildargs"`
	} `yaml:"images"`
	EnvironmentVariables map[string]string `yaml:"envvars"`
	BuildArguments       map[string]string `yaml:"buildargs"`
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
		EnvironmentVariables: siface.EnvironmentVariables,
		BuildArguments:       siface.BuildArguments,
	}

	s.Tree = NewTree()

	for k, v := range siface.Images {
		ba := make(map[string]string)
		for k2, v2 := range v.BuildArgs {
			switch sv := v2.(type) {
			case string:
				ba[k2] = sv
			default:
				return nil, fmt.Errorf("incorret value or type in buildargs (key: %s)", k2)
			}
		}

		i := NewImage(k, v.Tag, v.ContainerFile, v.Context, ba)

		s.Tree.Add(i, k)
	}

	return s, nil
}
