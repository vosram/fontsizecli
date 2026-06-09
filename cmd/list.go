package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "returns a list of basic ratios",
	Long: `This command displays a set of basic set of
common ratios a user can use to make their font size sets.`,
	Run: func(cmd *cobra.Command, args []string) {
		ratios := []string{
			"1.067 Minor Second",
			"1.125 Major Second",
			"1.2 Minor Third",
			"1.25 Major Third",
			"1.333 Perfect Fourth",
			"1.414 Augmented Fourth",
			"1.5 Perfect Fifth",
			"1.618 Golden Ratio",
		}
		fmt.Println("Common ratios in design:")
		for _, ratio := range ratios {
			fmt.Println("-", ratio)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
