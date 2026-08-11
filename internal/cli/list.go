package cli

import (
	"github.com/PavYtr/nofreedom/internal/units"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows supported units",
	Long:  "Shows supported units. You can use this command to see which units are available for conversion.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, categoryUnits := range units.RegistryByCategory {
			if len(categoryUnits) == 0 {
				continue
			}
			category := categoryUnits[0].Category
			cmd.Printf("[%s]:\n", category)
			for _, unit := range categoryUnits {
				cmd.Printf("\t %s (%s)\n", unit.Name, unit.Symbol)
			}
		}
		return nil
	},
}
