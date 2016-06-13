package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "euro2016",
	Short: "UEFA Euro 2016 cli",
	Long: "A client to follow UEFA Euro 2016.\n" +
		"Display competition schedule, live results, teams and rankings from the terminal.",
}
