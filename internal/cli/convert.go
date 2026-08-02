package cli

import (
	"fmt"
	"strconv"

	"github.com/PavYtr/nofreedom/internal/converter"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:     "convert <value> <from> <to>",
	Short:   "Convert units",
	Long:    "Convert units from one system to another. For example, you can convert meters to feet or kilograms to pounds.",
	Example: "nofreedom convert 10 m ft\nnofreedom convert 5 kg lb",
	Args:    cobra.ExactArgs(3),

	RunE: func(cmd *cobra.Command, args []string) error {
		value := args[0]
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid value %q: expected a number", args[0])
		}
		from := args[1]
		to := args[2]

		result, err := converter.Convert(val, from, to)
		if err != nil {
			return fmt.Errorf("convert %q to %q: %w", from, to, err)
		}
		cmd.Printf("%.2f\n", result)
		return nil
	},
}
