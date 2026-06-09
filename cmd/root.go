package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fontsizecli",
	Short: "Get a set of harmonius fontsizes for graphic design",
	Long: `Get a harmonius set of font sizes according to specific ratios
like: 1.067, 1.125, 1.2, 1.25, 1.333, 1.414, 1.5, 1.618

commands:
- web (creates font size sets for web design, h1-h6, p, etc...)
- custom (creates custom font size sets)
- list (shows list of common ratios)`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
