package cli

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/odpf/kay/config"
	"github.com/odpf/kay/internal/server"
	"github.com/odpf/kay/internal/store/postgres"
	"github.com/spf13/cobra"
)

func ServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server <command>",
		Aliases: []string{"s"},
		Short:   "Server management",
		Long:    "Server management commands.",
		Example: heredoc.Doc(`
			$ kay server start
			$ kay server start -c ./config.yaml
		`),
	}

	cmd.AddCommand(serverStartCommand())
	cmd.AddCommand(serverBootstrapCommand())

	return cmd
}

func serverStartCommand() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start kay server",
		Example: "kay server start",
		RunE: func(cmd *cobra.Command, args []string) error {
			appConfig, err := config.Load(configFile)
			if err != nil {
				panic(err)
			}

			return server.Start(appConfig)
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "", "Config file path")
	return cmd
}

func serverBootstrapCommand() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:     "bootstrap",
		Short:   "Bootstrap kay server",
		Example: "kay server bootstrap",
		RunE: func(cmd *cobra.Command, args []string) error {
			appConfig, err := config.Load(configFile)
			if err != nil {
				panic(err)
			}

			db, err := postgres.NewStore(appConfig.DB)

			if err != nil {
				panic(err)
			}
			return db.Bootstrap(cmd.Context())
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "", "Config file path")
	return cmd
}
