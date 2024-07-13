package shanghai

import (
	"github.com/dusansimic/shanghai"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", "shanghai.yaml", "Shangaifile to use")
	rootCmd.PersistentFlags().BoolVarP(&this, "this", "t", false, "work only on this image")
}

var filename string
var this bool
var logWriters shanghai.LogWriters

var rootCmd = &cobra.Command{
	Use: "shanghai",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	logWriters = shanghai.LogWriters{
		Err: rootCmd.ErrOrStderr(),
		Out: rootCmd.OutOrStdout(),
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
