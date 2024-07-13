package shanghai

import (
	"fmt"
	"os"

	"github.com/dusansimic/shanghai"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:               "push <image>",
	Short:             "Push hierarchies of container images",
	Args:              cobra.MinimumNArgs(1),
	ValidArgsFunction: imageCompletions,
	Run:               pushCommand,
}

func pushCommand(cmd *cobra.Command, args []string) {
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

	if err := s.Push(image); err != nil {
		fmt.Println(fmt.Errorf("failed to build image: %w", err))
		os.Exit(1)
	}
}
