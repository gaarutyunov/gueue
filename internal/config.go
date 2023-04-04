package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getConfig(cmd *cobra.Command, args []string) error {
	configPath, err := cmd.PersistentFlags().GetString("config")
	if err != nil {
		return err
	}

	viper.SetConfigFile(configPath)

	return viper.ReadInConfig()
}
