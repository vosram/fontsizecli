/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "returns a list of basic ratios",
	Long: `This command displays a set of basic set of
common ratios a user can use to make their font size sets.`,
	Run: func(cmd *cobra.Command, args []string) {
		ratios := []string{
			"1.067 Minor Second",
			"1.125 Major Second",
			"1.2 Minor Third (default)",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
