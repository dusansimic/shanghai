package file

import (
	"github.com/dusansimic/shanghai/polytree"
)

type Filer interface {
	Read(f string) (*File, error)
}

type File struct {
	Groups    map[string][]string
	Tree      polytree.PolyTree
	EnvVars   map[string]string
	BuildArgs map[string]string
}
