package cmd

import (
	"fmt"
	"todogolist/repository"
	service "todogolist/service/tasks"

	"github.com/spf13/cobra"
)

func buildDeleteTaskCmd(repository repository.Repository) *cobra.Command {

	opts := &runCommandOptions{}

	deleteTaskCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete all task for a given date or a given label",
		Run: func(cmd *cobra.Command, args []string) {
			resp := service.Delete(*&opts.TaskName, repository)
			fmt.Fprintf(cmd.OutOrStdout(), resp)

		},
	}

	f := deleteTaskCmd.Flags()
	f.StringVarP(&opts.TaskName, "task-name", "t", "", "Task Name we will like to delete")

	return deleteTaskCmd

}
