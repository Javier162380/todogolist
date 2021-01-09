package cmd

import (
	"fmt"
	"os"
	"todogolist/repository"

	"github.com/spf13/cobra"
)

// DefaultRepository Variable can be changed to choose the repository type wanted.
const DefaultRepository string = "ent"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd, err := buildRootCommand(DefaultRepository)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func buildRootCommand(repositoryType string) (*cobra.Command, error) {

	_, togoProfile := os.LookupEnv("TOGO_CONFIG")

	if togoProfile == false {
		os.Setenv("TOGO_CONFIG", fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".togo/config.yaml"))
	}

	repository := repository.NewRepository(repositoryType, os.Getenv("TOGO_CONFIG"))

	rootCmd := &cobra.Command{
		Use:   "todogolist",
		Short: "A useful CLI to get control of your daily life",
		Long: `The CLI to control your life.
You could easily use it create, update, get, list and delete your daily tasks.`,
	}

	rootCmd.AddCommand(buildCreateTaskCmd(repository))
	rootCmd.AddCommand(buildGetTaskCmd(repository))
	rootCmd.AddCommand(buildListTaskCmd(repository))
	rootCmd.AddCommand(buildSetUpCmd(repository))
	rootCmd.AddCommand(buildDeleteTaskCmd(repository))

	return rootCmd, nil

}
