package z

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/araddon/dateparse"
	"github.com/spf13/cobra"
)

var entryCmd = &cobra.Command{
	Use:   "entry ([flags]) [id]",
	Short: "Display or update activity",
	Long:  "Display or update tracked activity.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("%s %s", CharError, "Please provide a valid number")
			os.Exit(1)
		}

		entry, err := database.GetEntry(int64(id))
		if err != nil {
			fmt.Printf("%s %+v\n", CharError, err)
			os.Exit(1)
		}

		if begin != "" {
			entry.Begin, err = dateparse.ParseAny(begin)
			if err != nil {
				fmt.Printf("%s %+v\n", CharError, err)
				os.Exit(1)
			}
		}

		if finish != "" {
			entry.Finish, err = dateparse.ParseAny(finish)
			if err != nil {
				fmt.Printf("%s %+v\n", CharError, err)
				os.Exit(1)
			}
		}

		if project != "" {
			entry.Project = project
		}

		if task != "" {
			entry.Task = task
		}

		if notes != "" {
			entry.Notes = strings.ReplaceAll(notes, "\\n", "\n")
		}

		err = database.UpdateEntry(*entry)
		if err != nil {
			fmt.Printf("%s %+v\n", CharError, err)
			os.Exit(1)
		}
		fmt.Printf("%s %s\n", CharInfo, entry.GetOutput(true))
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
	entryCmd.Flags().StringVarP(&begin, "begin", "b", "", "Update date/time the activity began at")
	entryCmd.Flags().StringVarP(&finish, "finish", "s", "", "Update date/time the activity finished at")
	entryCmd.Flags().StringVarP(&project, "project", "p", "", "Update activity project")
	entryCmd.Flags().StringVarP(&notes, "notes", "n", "", "Update activity notes")
	entryCmd.Flags().StringVarP(&task, "task", "t", "", "Update activity task")
	entryCmd.Flags().BoolVar(&fractional, "decimal", true, "Show fractional hours in decimal format instead of minutes")

	var err error
	database, err = InitDB()
	if err != nil {
		fmt.Printf("%s %+v\n", CharError, err)
		os.Exit(1)
	}
}
