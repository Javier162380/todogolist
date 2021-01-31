package cmd

import (
	"fmt"
	"todogolist/repository"
	service "todogolist/service/tasks"

	"github.com/spf13/cobra"
)

func buildListTaskCmd(repository repository.Repository) *cobra.Command {

	opts := &runCommandOptions{}

	listTaskCmd := &cobra.Command{
		Use:   "list",
		Short: "List all task for a given date or a given label",
		Run: func(cmd *cobra.Command, args []string) {
			resp := service.List(*&opts.Limit, repository)

			fmt.Fprintf(cmd.OutOrStdout(), resp)

		},
	}

	f := listTaskCmd.Flags()
	f.IntVarP(&opts.Limit, "limit", "L", 20, "Limit Results fetch by the query")

	return listTaskCmd

}
