package shanghai

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ListImages(f string) ([]string, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var file Shanghaifile
	if err := yaml.Unmarshal(b, &file); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}

	keys := make([]string, len(file.Images))
	for k := range file.Images {
		keys = append(keys, k)
	}

	return keys, nil
}
