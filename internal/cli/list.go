package cli

import (
	"github.com/PavYtr/nofreedom/internal/units"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows supported units",
	Long:  "Shows supported units. You can use this command to see which units are available for conversion.",
	RunE: func(cmd *cobra.Command, args []string) error {
		for category, unit := range units.RegistryByCategory {
			cmd.Printf("[%s]:\n", category)
			for _, u := range unit {
				cmd.Printf("\t %s (%s)\n", u.Name, u.Symbol)
			}
		}
		return nil
	},
}
