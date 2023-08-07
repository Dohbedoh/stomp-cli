package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCheckCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "check",
		Short: "Check and Test STOMP",
		Long:  "Check and Test STOMP",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Test STOMP connections")
		},
	}
	command.AddCommand(NewCheckConnectionCommand())
	return command
}
