package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initializationCmd represents the initialization command
var initializationCmd = &cobra.Command{
	Use:     "initialization",
	Aliases: []string{"Initialize", "init"},
	Short:   "A brief description of your command",
	Long:    `A longer description of your command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initialization called")
	},
}

func init() {
	rootCmd.AddCommand(initializationCmd)
}
