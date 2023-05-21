package cmd

import (
	"fmt"
	"os"
	"shanghai/libshg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringVarP(&file, "file", "f", "Shangaifile", "Shangaifile to use")
	buildCmd.MarkFlagFilename("file", "shg", "yaml", "yml")
	buildCmd.Flags().BoolVarP(&check, "check", "c", false, "check Shangaifile for errors")
}

var file string
var image string
var check bool

var buildCmd = &cobra.Command{
	Use:   "build <image>",
	Short: "Build hierarchies of container images",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		l, err := libshg.ListImages(file)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return l, cobra.ShellCompDirectiveNoFileComp
	},
	Run: command,
}

func command(cmd *cobra.Command, args []string) {
	image = args[0]

	file, err := libshg.ReadShanghaifile(file)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read Shangaifile: %w", err))
		os.Exit(1)
	}

	if err := libshg.ValidateShanghaifile(file, image); err != nil {
		fmt.Println(fmt.Errorf("failed to validate Shangaifile: %w", err))
		os.Exit(1)
	}

	if check {
		return
	}

	cfile, err := libshg.SearchConfigFile()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to find config file: %w", err))
		os.Exit(1)
	}

	c, err := libshg.ReadConfig(cfile)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read config file: %w", err))
		os.Exit(1)
	}

	if err := libshg.BuildImages(c, file, image); err != nil {
		fmt.Println(fmt.Errorf("failed to build image: %w", err))
		os.Exit(1)
	}
}
