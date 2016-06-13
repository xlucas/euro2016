package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(teamsCmd)
}

var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Show teams being part of the competition.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
