package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nofreedom",
	Short: "Unit converter with CLI interface",
	Long:  "A unit converter that embraces the glorious superiority of the metric system.",
}

func init() {
	rootCmd.AddCommand(convertCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
