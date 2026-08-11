package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "nofreedom",
		Short:         "Unit converter with CLI interface",
		Long:          "A unit converter that embraces the glorious superiority of the metric system.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(listCmd)

	return rootCmd
}

func Execute() error {
	return NewRootCmd().Execute()
}
