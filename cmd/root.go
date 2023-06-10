package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "shanghi.yaml", "Shangaifile to use")
}

var file string

var rootCmd = &cobra.Command{
	Use: "shanghai",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
