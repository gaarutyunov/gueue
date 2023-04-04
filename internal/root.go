package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "gueue",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func Run() {
	log.SetFormatter(&log.JSONFormatter{})

	die := make(chan interface{})
	defer func() {
		close(die)
	}()

	go func() {
		err := rootCmd.Execute()
		die <- err
	}()

	reason := <-die

	switch reason := reason.(type) {
	case error:
		log.Fatal(reason)
	default:
		return
	}
}
