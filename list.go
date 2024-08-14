package shanghai

import (
	"fmt"

	yamlfile "github.com/dusansimic/shanghai/file/yaml"
)

func ListImages(f string) ([]string, error) {
	fileStruct, err := yamlfile.New().Read(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	is := fileStruct.Tree.Nodes()
	keys := make([]string, len(is))
	for _, im := range is {
		keys = append(keys, im.Name())
	}

	return keys, nil
}

func ListGroups(f string) ([]string, error) {
	fileStruct, err := yamlfile.New().Read(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	gs := make([]string, len(fileStruct.Groups))
	i := 0
	for k := range fileStruct.Groups {
		gs[i] = k
		i++
	}

	return gs, nil
}
