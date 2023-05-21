package cmd

import (
	"fmt"
	"os"
	"shanghai/libshg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate <image>",
	Short: "Validate Shanghaifile",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		l, err := libshg.ListImages(file)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return l, cobra.ShellCompDirectiveNoFileComp
	},
	Run: validateCommand,
}

func validateCommand(cmd *cobra.Command, args []string) {
	image := args[0]

	file, err := libshg.ReadShanghaifile(file)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read Shanghaifile: %w", err))
		os.Exit(1)
	}

	if err := libshg.ValidateShanghaifile(file, image); err != nil {
		fmt.Println(fmt.Errorf("failed to validate Shangaifile: %w", err))
		os.Exit(1)
	}
}
