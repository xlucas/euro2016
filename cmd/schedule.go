package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
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
	scheduleCmd.AddCommand(fullScheduleCmd)
	scheduleCmd.AddCommand(statusScheduleCmd)
	scheduleCmd.AddCommand(todayScheduleCmd)
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Print the competition schedule",
}

var fullScheduleCmd = &cobra.Command{
	Use:   "full",
	Short: "Print full schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			start, _ = time.Parse(time.RFC3339, "2016-06-10T00:00:00Z")
			end, _   = time.Parse(time.RFC3339, "2016-07-11T00:00:00Z")
		)
		return showFixtures(start, end, "")
	},
}

var todayScheduleCmd = &cobra.Command{
	Use:   "today",
	Short: "Print today's schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			now      = time.Now()
			nextDay  = now.Add(24 * time.Hour)
			start, _ = time.Parse(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00Z", now.Year(), now.Month(), now.Day()))
			end, _   = time.Parse(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00Z", nextDay.Year(), nextDay.Month(), nextDay.Day()))
		)
		return showFixtures(start, end, "")
	},
}

var statusScheduleCmd = &cobra.Command{
	Use:   "status [FINISHED|IN_PLAY|TIMED]",
	Short: "Print schedule for games with a specific status",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			start, _ = time.Parse(time.RFC3339, "2016-06-10T00:00:00Z")
			end, _   = time.Parse(time.RFC3339, "2016-07-11T00:00:00Z")
		)
		if len(args) != 1 {
			return fmt.Errorf("Expecting game state as argument")
		}
		return showFixtures(start, end, args[0])
	},
}

func showFixtures(start, end time.Time, status string) error {
	f, err := getFixtures(start, end, status)
	if err != nil {
		return err
	}
	printFixtures(f, os.Stdout)

	return nil
}

func getFixtures(from, to time.Time, status string) ([]Fixture, error) {
	var (
		schedule Schedule
		fixtures []Fixture
	)

	err := client.Get("/fixtures", &schedule)
	if err != nil {
		return nil, err
	}

	for _, f := range schedule.Fixtures {
		if !f.Date.After(to) && !f.Date.Before(from) {
			if f.Status == status || status == "" {
				fixtures = append(fixtures, f)
			}
		}
	}

	return fixtures, nil
}

func printFixtures(fixtures []Fixture, out *os.File) {
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Day", "Home Team", "Score", "Away Team", "Status", "Date"})

	for _, f := range fixtures {
		data := []string{
			fmt.Sprintf("%d", f.Matchday),
			f.HomeTeam,
			fmt.Sprintf("%d - %d", f.Result.GoalsHome, f.Result.GoalsAway),
			f.AwayTeam,
			f.Status,
			f.Date.Format(time.RFC822),
		}
		table.Append(data)
	}
	table.Render()
}
