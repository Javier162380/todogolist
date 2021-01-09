package cmd

import (
	"fmt"
	"todogolist/repository"
	service "todogolist/service/tasks"

	"github.com/spf13/cobra"
)

func buildGetTaskCmd(repository repository.Repository) *cobra.Command {

	opts := &runOptions{}

	getTaskCmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve tasks",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := service.Get(*&opts.TaskName, *&opts.TaskLabel, *&opts.TaskDate, *&opts.TaskEndDate, *&opts.Limit, repository)

			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), fmt.Sprintf("%s", err))
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), resp)
			}
		},
	}

	f := getTaskCmd.Flags()
	f.StringVarP(&opts.TaskName, "task-name", "t", "", "Name of the task")
	f.StringVarP(&opts.TaskLabel, "task-label", "l", "", `Lable of the task. 
Possible Values: unknown|tech|friends|sport|work|reading|family`)
	f.StringVarP(&opts.TaskDate, "start-date", "s", "", "Start Date/Time of the task")
	f.StringVarP(&opts.TaskEndDate, "end-date", "e", "", "End Date/Time of the task")
	f.IntVarP(&opts.Limit, "limit", "L", 0, "Limit values to retrieve")

	return getTaskCmd

}
