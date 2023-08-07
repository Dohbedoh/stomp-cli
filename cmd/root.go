package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stomp",
	Short: "STOMP CLI",
	Long:  "STOMP Command Line Utility",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test STOMP connections")
	},
}

func init() {
	rootCmd.AddCommand(NewCheckCommand())
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
