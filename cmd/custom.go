package cmd

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vosram/fontsizecli/internal/utils"
)

var customCmd = &cobra.Command{
	Use:   "custom <\"h\" | \"l\"> <base-font-size> <ratio> <num-of-steps>",
	Short: "Create a set of font sizes based on a ratio",
	Long: `Create a set of font sizes based on a ratio and according to a base
font size either as the base being the height value or the lowest value.`,
	Args: cobra.MatchAll(cobra.ExactArgs(4)),
	RunE: func(cmd *cobra.Command, args []string) error {
		outputName, err := cmd.Flags().GetString("output")
		if err != nil {
			return nil
		}

		baseStart := strings.ToLower(args[0])
		if baseStart != "h" && baseStart != "l" {
			return errors.New("First argument should be either h or l")
		}

		baseSize, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		ratio, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return err
		}
		steps, err := strconv.Atoi(args[3])
		if err != nil {
			return err
		}

		if baseStart == "h" {
			lastValue := baseSize
			content := ""
			for i := 1; i <= steps; i++ {
				if i == 1 {
					content += fmt.Sprintf("%d - %dpx\n", i, baseSize)
					continue
				}
				result := int(math.Round(float64(lastValue) / ratio))
				content += fmt.Sprintf("%d - %dpx\n", i, result)
				lastValue = result
			}
			fmt.Println(content)
			if outputName != "" {
				err := utils.SaveResultToFile(outputName, content)
				if err != nil {
					return err
				}
			}
			return nil
		}

		// at this point baseStart is l
		lastValue := baseSize
		content := ""
		for i := 1; i <= steps; i++ {
			if i == 1 {
				content += fmt.Sprintf("%d - %dpx\n", i, baseSize)
				continue
			}
			result := int(math.Round(float64(lastValue) * ratio))
			content += fmt.Sprintf("%d - %dpx\n", i, result)
			lastValue = result
		}
		fmt.Println(content)
		if outputName != "" {
			err := utils.SaveResultToFile(outputName, content)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(customCmd)

	customCmd.Flags().StringP("output", "o", "", "filename to save the output to, .txt is added automatically")
}
