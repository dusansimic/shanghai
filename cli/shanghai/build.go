package shanghai

import (
	"fmt"
	"os"
	"shanghai"

	"github.com/spf13/cobra"
)

func init() {
	buildCmd.PersistentFlags().BoolVarP(&buildThis, "this", "t", false, "Build only this image")
	rootCmd.AddCommand(buildCmd)
}

var buildThis bool

var buildCmd = &cobra.Command{
	Use:   "build [options] <image>",
	Short: "Build hierarchies of container images",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		l, err := shanghai.ListImages(file)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return l, cobra.ShellCompDirectiveNoFileComp
	},
	Run: buildCommand,
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

	s := shanghai.NewSession(cfg, shg, buildThis, logWriters)

	if err := s.Build(image); err != nil {
		fmt.Println(fmt.Errorf("failed to build image: %w", err))
		os.Exit(1)
	}
}
