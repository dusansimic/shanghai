package shanghai

import (
	"fmt"
	"shanghai"
)

func readConfig() (*shanghai.Config, error) {
	cfile, err := shanghai.SearchConfigFile()
	if err != nil {
		return nil, fmt.Errorf("failed to find config file: %w", err)
	}

	c, err := shanghai.ReadConfig(cfile)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func readShanghaifile(image string) (*shanghai.Shanghaifile, error) {
	file, err := shanghai.ReadShanghaifile(file)
	if err != nil {
		return nil, err
	}

	return file, nil
}
