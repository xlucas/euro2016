package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/xlucas/euro2016/util"
)

const (
	endpoint = "http://api.football-data.org/v1/soccerseasons/424"
)

type Schedule struct {
	Fixtures []Fixture `json:"fixtures"`
}

type Fixture struct {
	AwayTeam string    `json:"awayTeamName"`
	Date     time.Time `json:"date"`
	HomeTeam string    `json:"homeTeamName"`
	Matchday uint8     `json:"matchday"`
	Result   Result    `json:"result"`
	Status   string    `json:"status"`
}

type Result struct {
	GoalsAway uint8 `json:"goalsAwayTeam"`
	GoalsHome uint8 `json:"goalsHomeTeam"`
}

func init() {
	RootCmd.AddCommand(scheduleCmd)
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Print the competition schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			c *util.JSONClient
			s Schedule
		)

		c = util.NewJSONClient(endpoint, "")
		err := c.Get("/fixtures", &s)
		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Home Team", "  ", "  ", "Away Team", "Status", "Date"})

		for _, f := range s.Fixtures {
			data := []string{
				f.HomeTeam,
				fmt.Sprintf("%d", f.Result.GoalsHome),
				fmt.Sprintf("%d", f.Result.GoalsAway),
				f.AwayTeam,
				f.Status,
				f.Date.Format("Jan _2 15:04"),
			}
			table.Append(data)
		}

		table.Render()

		return nil
	},
}
