package cmd

import "github.com/spf13/cobra"

const (
	endpoint = "http://api.football-data.org/v1/soccerseasons/424"
)

var RootCmd = &cobra.Command{
	Use:   "euro2016",
	Short: "UEFA Euro 2016 CLI",
	Long: "A client to follow UEFA Euro 2016.\n" +
		"Display competition schedule, live results, teams and rankings from the terminal.",
}
