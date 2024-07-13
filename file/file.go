package file

import (
	"fmt"
	"os"

	"github.com/dusansimic/shanghai/image"
	"github.com/dusansimic/shanghai/polytree"
	"gopkg.in/yaml.v3"
)

type File struct {
	Tree      polytree.PolyTree
	EnvVars   map[string]string
	BuildArgs map[string]string
}

type file struct {
	Images map[string]struct {
		Tag           string
		Tags          []string
		ContainerFile string
		Context       string
		BuildArgs     map[string]interface{}
		Parents       []string
	}
	EnvVars   map[string]string
	BuildArgs map[string]string
}

func Read(f string) (*File, error) {
	d, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read Shanghaifile: %w", err)
	}

	file := &file{}
	if yaml.Unmarshal(d, &file) != nil {
		return nil, fmt.Errorf("failed to unmarshal Shanghaifile: %w", err)
	}

	s := &File{
		EnvVars:   file.EnvVars,
		BuildArgs: file.BuildArgs,
	}

	s.Tree = polytree.NewPolyTree()

	for k, v := range file.Images {
		ba := make(map[string]string)
		for k2, v2 := range v.BuildArgs {
			switch sv := v2.(type) {
			case string:
				ba[k2] = sv
			case int:
				ba[k2] = fmt.Sprint(sv)
			default:
				return nil, fmt.Errorf("incorret value or type in buildargs (key: %s)", k2)
			}
		}

		if v.Tag == "" && len(v.Tags) == 0 {
			return nil, fmt.Errorf("tag or tags not specified")
		}

		if len(v.Tags) == 0 {
			v.Tags = []string{v.Tag}
		}

		i := image.NewImage(k, v.Tags, v.ContainerFile, v.Context, ba, v.Parents)

		if err := s.Tree.Add(i); err != nil {
			return nil, fmt.Errorf("duplicate image in tree '%s'", i.Name())
		}
	}

	return s, nil
}
