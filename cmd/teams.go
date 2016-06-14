package cmd

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(teamsCmd)
}

type LineUp struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	Name string `json:"name"`
}

var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Show teams being part of the competition.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			lineup LineUp
			names  []string
		)

		err := client.Get("/teams", &lineup)
		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Team"})

		for _, k := range lineup.Teams {
			names = append(names, k.Name)
		}

		sort.Strings(names)

		for _, name := range names {
			table.Append([]string{name})
		}

		table.Render()

		return nil
	},
}
