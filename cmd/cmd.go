package cmd

import (
	"os"

	addevents "Priofirst-CLI/cmd/add"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "PrioFirstCLI",
	Short: "prfs",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addevents.NewAddCmd())
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
