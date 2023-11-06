package shanghai

import (
	"fmt"
	"os"
	"shanghai"

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
		l, err := shanghai.ListImages(file)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return l, cobra.ShellCompDirectiveNoFileComp
	},
	Run: validateCommand,
}

func validateCommand(cmd *cobra.Command, args []string) {
	image := args[0]

	file, err := shanghai.ReadShanghaifile(file)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read Shanghaifile: %w", err))
		os.Exit(1)
	}

	s := shanghai.NewSession(nil, file, logWriters)

	if err := s.ValidateShanghaifile(image); err != nil {
		fmt.Println(fmt.Errorf("failed to validate image: %w", err))
		os.Exit(1)
	}
}
