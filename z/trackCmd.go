package z

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Tracking time",
	Long:  "Track new activity, which can either be kept running until 'finish' is being called or parameterized to be a finished activity.",
	Run: func(cmd *cobra.Command, args []string) {

		entry, err := database.GetRunningEntry()
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				//ignore this, as there are no tasks running, which is what we want.
			default:
				fmt.Printf("%s %+v\n", CharError, err)
				os.Exit(1)
			}
		}
		if entry != nil {
			fmt.Printf("%s A task is already running, you have to finish it first before you start a new one.\n\nType 'zeit finish' to do so.", CharError)
			os.Exit(1)
		}
		if task == "" {
			fmt.Printf("%s Can not track empty task.\nPlease assign a task via --task to track\n", CharError)
		}
		if project == "" {
			fmt.Printf("%s Can not track empty project.\nPlease assign a project via --project\n", CharError)
			os.Exit(1)
		}
		newEntry := NewEntry(project, task)
		if notes != "" {
			newEntry.Notes = notes
		}
		err = database.AddEntry(&newEntry, true)
		if err != nil {
			fmt.Printf("something went wrong. Error: %s", err.Error())
			os.Exit(1)
		}

		fmt.Print(newEntry.GetStartTrackingStr())
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
	trackCmd.Flags().StringVarP(&begin, "begin", "b", "", "Time the activity should begin at\n\nEither in the formats 16:00 / 4:00PM \nor relative to the current time, \ne.g. -0:15 (now minus 15 minutes), +1.50 (now plus 1:30h).")
	trackCmd.Flags().StringVarP(&finish, "finish", "s", "", "Time the activity should finish at\n\nEither in the formats 16:00 / 4:00PM \nor relative to the current time, \ne.g. -0:15 (now minus 15 minutes), +1.50 (now plus 1:30h).\nMust be after --begin time.")
	trackCmd.Flags().StringVarP(&project, "project", "p", "", "Project to be assigned")
	trackCmd.Flags().StringVarP(&task, "task", "t", "", "Task to be assigned")
	trackCmd.Flags().StringVarP(&notes, "notes", "n", "", "Activity notes")

	var err error
	database, err = InitDB()
	if err != nil {
		fmt.Printf("%s %+v\n", CharError, err)
		os.Exit(1)
	}
}
