package cmd

import (
	"fmt"
	"log"
	"todogolist/repository"
	service "todogolist/service/tasks"

	"github.com/spf13/cobra"
)

func buildSetUpCmd(repository repository.Repository) *cobra.Command {

	opts := &runOptions{}

	setUpCmd := &cobra.Command{
		Use:   "setup",
		Short: "Set up/update all the needed DB model",
		Long: `This command create/updated all the needed DB objects, tables, etc.        
        This command needs to be the first command executed before creating the task`,
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := service.Setup(*&opts.dryRun, repository)

			if err != nil {
				log.Fatalf("%s", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), resp)

		},
	}

	f := setUpCmd.Flags()
	f.BoolVarP(&opts.dryRun, "dry-run", "d", false, "Dry Run retrieves the migration previous exeuction")

	return setUpCmd
}
