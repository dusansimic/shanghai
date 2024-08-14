package yaml

import (
	"fmt"
	"os"

	"github.com/dusansimic/shanghai/file"
	"github.com/dusansimic/shanghai/image"
	"github.com/dusansimic/shanghai/polytree"
	"gopkg.in/yaml.v3"
)

type filer struct{}

type filestruct struct {
	Groups map[string][]string
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

func New() file.Filer {
	return &filer{}
}

func (filer) Read(f string) (*file.File, error) {
	d, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read Shanghaifile: %w", err)
	}

	filestruct := &filestruct{}
	if yaml.Unmarshal(d, &filestruct) != nil {
		return nil, fmt.Errorf("failed to unmarshal Shanghaifile: %w", err)
	}

	s := &file.File{
		Groups:    filestruct.Groups,
		EnvVars:   filestruct.EnvVars,
		BuildArgs: filestruct.BuildArgs,
	}

	s.Tree = polytree.NewPolyTree()

	for k, v := range filestruct.Images {
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
