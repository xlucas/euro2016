package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type League struct {
	Day       int8                  `json:"matchday"`
	Name      string                `json:"leagueCaption"`
	Standings map[string][]TeamRank `json:"standings"`
}

type TeamRank struct {
	GoalAverage  int8   `json:"goalDifference"`
	GoalsFor     int8   `json:"goals"`
	GoalsAgainst int8   `json:"goalsAgainst"`
	Group        string `json:"group"`
	Played       int8   `json:"playedGames"`
	Points       int8   `json:"points"`
	Rank         int8   `json:"rank"`
	Team         string `json:"team"`
}

func init() {
	RootCmd.AddCommand(rankingsCmd)
	rankingsCmd.AddCommand(fullRankingsCmd)
	rankingsCmd.AddCommand(groupRankingsCmd)
}

var rankingsCmd = &cobra.Command{
	Use:   "rankings",
	Short: "Show group rankings",
}

var fullRankingsCmd = &cobra.Command{
	Use:   "full",
	Short: "Show rankings for all groups",
	RunE: func(cmd *cobra.Command, args []string) error {
		var names []string

		groups, err := getGroups()
		if err != nil {
			return err
		}
		for k := range groups {
			names = append(names, k)
		}

		sort.Strings(names)
		table := tablewriter.NewWriter(os.Stdout)
		printHeader(table)

		for i, name := range names {
			printGroup(table, name, groups[name], i == len(names)-1)
		}

		table.Render()

		return nil
	},
}

var groupRankingsCmd = &cobra.Command{
	Use:   "group [A|B|C|D|E|F]",
	Short: "Show rankings for a group",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Expecting group name as argument")
		}

		group := args[0]
		groups, err := getGroups()
		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		printHeader(table)
		printGroup(table, group, groups[group], true)
		table.Render()

		return nil
	},
}

func getGroups() (map[string][]TeamRank, error) {
	var league League
	err := client.Get("/leagueTable", &league)
	return league.Standings, err
}

func printHeader(table *tablewriter.Table) {
	table.SetHeader([]string{"Group", "POS", "Team", "P", "GF", "GA", "GD", "PTS"})
}

func printGroup(table *tablewriter.Table, name string, ranks []TeamRank, last bool) {
	for _, rank := range ranks {
		table.Append([]string{
			rank.Group,
			fmt.Sprintf("%d", rank.Rank),
			rank.Team,
			fmt.Sprintf("%d", rank.Played),
			fmt.Sprintf("%d", rank.GoalsFor),
			fmt.Sprintf("%d", rank.GoalsAgainst),
			fmt.Sprintf("%d", rank.GoalAverage),
			fmt.Sprintf("%d", rank.Points),
		})
	}
	if !last {
		table.Append([]string{"", "", "", "", "", "", "", ""})
	}
}
