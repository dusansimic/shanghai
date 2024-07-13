package shanghai

import (
	"fmt"

	"github.com/dusansimic/shanghai"
	"github.com/dusansimic/shanghai/file"
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

func readShanghaifile(image string) (*file.File, error) {
	filestruct, err := file.Read(filename)
	if err != nil {
		return nil, err
	}

	return filestruct, nil
}
