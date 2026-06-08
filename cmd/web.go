/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vosram/fontsizecli/internal/utils"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web <base-size> <ratio>",
	Short: "Create a set of fonts for h1-h6, p, smp, and xsmp",
	Long:  `Create a set of fonts for h1-h6, p, two extra p sizes "sm" and xsm"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseSize, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		deskRatio, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			return err
		}
		mobileBase, err := cmd.Flags().GetInt("mob-base")
		if err != nil {
			return err
		}
		mobileRatio, err := cmd.Flags().GetFloat64("mob-ratio")
		if err != nil {
			return err
		}
		outputName, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		if mobileBase == 0 && mobileRatio == 0 {
			// Mobile size and ratio not set
			values := utils.CreateWebSizeValues(baseSize, deskRatio)
			content := fmt.Sprintln("Desktop/Mobile Sizes:") +
				utils.FormatWebSizes(values)
			fmt.Println(content)
			if outputName != "" {
				err := utils.SaveResultToFile(outputName, content)
				if err != nil {
					return err
				}
			}
			return nil
		}

		if mobileBase == 0 {
			mobileBase = baseSize
		}
		if mobileRatio == 0 {
			mobileRatio = deskRatio
		}
		deskValues := utils.CreateWebSizeValues(baseSize, deskRatio)
		mobValues := utils.CreateWebSizeValues(mobileBase, mobileRatio)
		content := fmt.Sprintln("Desktop Sizes:") +
			utils.FormatWebSizes(deskValues) +
			fmt.Sprintln("\nMobile Sizes:") +
			utils.FormatWebSizes(mobValues)

		fmt.Println(content)
		if outputName != "" {
			err := utils.SaveResultToFile(outputName, content)
			if err != nil {
				return err
			}
		}
		return nil
	},
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	webCmd.Flags().StringP("output", "o", "", "Name for output file. Extension .txt is automatically added")
	webCmd.Flags().Float64("mob-ratio", 0, "Ratio for mobile sizes")
	webCmd.Flags().Int("mob-base", 0, "Mobile base font size")
}
