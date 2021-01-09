package cmd

import (
	"fmt"
	"todogolist/repository"
	service "todogolist/service/tasks"

	"github.com/spf13/cobra"
)

var (
	taskName         string
	taskLabel        string
	taskDate         string
	taskTimeInvested int
	taskPeriodicity  string
)

func buildCreateTaskCmd(repository repository.Repository) *cobra.Command {

	opts := &runOptions{}
	createTaskCmd := &cobra.Command{
		Use:   "create",
		Short: "Crete a new task",
		Long:  `Create a new task in the togolist`,
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := service.Create(*&opts.TaskName, *&opts.TaskLabel, *&opts.TaskDate, *&opts.debugMode, *&opts.TaskTimeInvested, *&opts.TaskPeriodicity, repository)

			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", err))
			}

			fmt.Fprintf(cmd.OutOrStdout(), resp)
		},
	}

	f := createTaskCmd.Flags()
	f.StringVarP(&opts.TaskName, "task-name", "t", "", "Name of the task")
	f.StringVarP(&opts.TaskLabel, "task-label", "l", "", `Label of the task. 
Possible Values: unknown|tech|friends|sport|work|reading|family`)
	f.StringVarP(&opts.TaskDate, "task-date", "D", "", "Date/Time of the task")
	f.IntVarP(&opts.TaskTimeInvested, "tasktime-invested", "i", 0, "Time you are going to invest in the task in seconds")
	f.StringVarP(&opts.TaskPeriodicity, "task-periodicity", "p", "", `Periodicity of the task.
Possible Values: (w|weekly|d|daily|m|monthly`)
	f.BoolVarP(&opts.debugMode, "verbose", "v", false, "Verbose output")

	return createTaskCmd
}
