package internal

import (
	"github.com/gaarutyunov/gueue/pkg/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:               "server",
	RunE:              serve,
	PersistentPreRunE: getConfig,
}

func init() {
	flags := serverCmd.PersistentFlags()

	flags.StringP("config", "c", "gueue.yaml", "Config path")
}

func serve(cmd *cobra.Command, args []string) error {
	srv := server.NewServer()

	return srv.Start(cmd.Context())
}
