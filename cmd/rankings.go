package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(rankingsCmd)
}

var rankingsCmd = &cobra.Command{
	Use:   "rankings",
	Short: "Show group rankings",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
