package shanghai

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Shanghaifile struct {
	Tree                 PolyTree
	EnvironmentVariables map[string]string
	BuildArguments       map[string]string
}

type shanghaifile struct {
	Images map[string]struct {
		Tag           string
		ContainerFile string
		Context       string
		BuildArgs     map[string]interface{}
		Parents       []string
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

	s.Tree = NewPolyTree()

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

		i := NewImage(k, v.Tag, v.ContainerFile, v.Context, ba, v.Parents)

		if err := s.Tree.Add(i); err != nil {
			return nil, fmt.Errorf("duplicate image in tree '%s'", i.Name())
		}
	}

	return s, nil
}
