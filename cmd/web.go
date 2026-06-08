/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"os"
	"path"
	"strconv"

	"github.com/spf13/cobra"
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
			values := createWebSizeValues(baseSize, deskRatio)
			content := fmt.Sprintln("Desktop/Mobile Sizes:") +
				formatWebSizes(values)
			fmt.Println(content)
			if outputName != "" {
				err := saveResultToFile(outputName, content)
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
		deskValues := createWebSizeValues(baseSize, deskRatio)
		mobValues := createWebSizeValues(mobileBase, mobileRatio)
		content := fmt.Sprintln("Desktop Sizes:") +
			formatWebSizes(deskValues) +
			fmt.Sprintln("\nMobile Sizes:") +
			formatWebSizes(mobValues)

		fmt.Println(content)
		if outputName != "" {
			err := saveResultToFile(outputName, content)
			if err != nil {
				return err
			}
		}
		return nil
	},
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
}

type sizeValueSet struct {
	h1  int
	h2  int
	h3  int
	h4  int
	h5  int
	h6  int
	p   int
	smp int
	xsp int
}

func createWebSizeValues(baseSize int, ratio float64) sizeValueSet {
	highValues := make([]int, 0, 6)
	lastValue := baseSize
	for i := 0; i < 6; i++ {
		value := float64(lastValue) * ratio
		result := int(math.Round(value))
		highValues = append(highValues, result)
		lastValue = result
	}

	lowValues := make([]int, 0, 3)
	lastValue = baseSize
	for i := 0; i < 2; i++ {
		value := float64(lastValue) / ratio
		result := int(math.Round(value))
		lowValues = append(lowValues, result)
		lastValue = result
	}

	values := sizeValueSet{
		h1:  highValues[5],
		h2:  highValues[4],
		h3:  highValues[3],
		h4:  highValues[2],
		h5:  highValues[1],
		h6:  highValues[0],
		p:   baseSize,
		smp: lowValues[0],
		xsp: lowValues[1],
	}
	return values
}

func formatWebSizes(values sizeValueSet) string {
	content := fmt.Sprintf("- %s: %dpx\n", " h1", values.h1) +
		fmt.Sprintf("- %s: %dpx\n", " h2", values.h2) +
		fmt.Sprintf("- %s: %dpx\n", " h3", values.h3) +
		fmt.Sprintf("- %s: %dpx\n", " h4", values.h4) +
		fmt.Sprintf("- %s: %dpx\n", " h5", values.h5) +
		fmt.Sprintf("- %s: %dpx\n", " h6", values.h6) +
		fmt.Sprintf("- %s: %dpx\n", "  p", values.p) +
		fmt.Sprintf("- %s: %dpx\n", "smp", values.smp) +
		fmt.Sprintf("- %s: %dpx\n", "xsp", values.xsp)
	return content
}

func saveResultToFile(filename, content string) error {
	filename = filename + ".txt"
	file, err := os.Create(path.Join(".", filename))
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
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
