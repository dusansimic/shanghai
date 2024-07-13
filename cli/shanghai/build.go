package shanghai

import (
	"fmt"
	"os"

	"github.com/dusansimic/shanghai"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:               "build [options] <image>",
	Short:             "Build hierarchies of container images",
	Args:              cobra.MinimumNArgs(1),
	ValidArgsFunction: imageCompletions,
	Run:               buildCommand,
}

func buildCommand(cmd *cobra.Command, args []string) {
	image := args[0]

	shg, err := readShanghaifile(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg, err := readConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := shanghai.NewSession(cfg, shg, this, group, logWriters)

	if err := s.Build(image); err != nil {
		fmt.Println(fmt.Errorf("failed to build image: %w", err))
		os.Exit(1)
	}
}
