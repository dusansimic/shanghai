package shanghai

import (
	"fmt"
	"os"
	"shanghai"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push <image>",
	Short: "Push hierarchies of container images",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		l, err := shanghai.ListImages(file)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return l, cobra.ShellCompDirectiveNoFileComp
	},
	Run: pushCommand,
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

	if err := shanghai.PushImages(cfg, shg, logWriters, image); err != nil {
		fmt.Println(fmt.Errorf("failed to push image: %w", err))
		os.Exit(1)
	}
}
