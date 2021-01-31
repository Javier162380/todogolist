package cmd

import (
	api "todogolist/api/infrastructure"
	"todogolist/repository"

	"github.com/spf13/cobra"
)

func buildStartAPICmd(repository repository.Repository) *cobra.Command {

	opts := &runAPIOptions{}
	startAPICmd := &cobra.Command{
		Use:   "startapi",
		Short: "Start a new api server",
		Long:  `Start a new api server with all the `,
		Run: func(cmd *cobra.Command, args []string) {

			srv := api.NewServer(opts.host, opts.port, repository)
			srv.Run()

		},
	}

	f := startAPICmd.Flags()
	f.UintVarP(&opts.port, "port", "p", 8000, "Port where the API it is going to be deployed")
	f.StringVarP(&opts.host, "host", "H", "", "Host where the API it is going to be serve")

	return startAPICmd
}
