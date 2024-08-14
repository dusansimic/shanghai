package shanghai

import (
	"fmt"

	"github.com/dusansimic/shanghai"
	"github.com/dusansimic/shanghai/file"
	yamlfile "github.com/dusansimic/shanghai/file/yaml"
	"github.com/spf13/cobra"
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
	filestruct, err := yamlfile.New().Read(filename)
	if err != nil {
		return nil, err
	}

	return filestruct, nil
}

func imageCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var l []string
	var err error
	if group {
		l, err = shanghai.ListGroups(filename)
	} else {
		l, err = shanghai.ListImages(filename)
	}

	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return l, cobra.ShellCompDirectiveNoFileComp
}
