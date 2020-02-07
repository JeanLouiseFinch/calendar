package httpserver

import (
	"errors"

	"github.com/JeanLouiseFinch/calendar/config"
	"github.com/JeanLouiseFinch/calendar/logger"
	"github.com/JeanLouiseFinch/calendar/server/http"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "http server",
		Short: "http server run cmd",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfgFile != "" {
				cfg, err := config.GetConfig(cfgFile)
				if err != nil {
					return err
				}
				log, err := logger.GetLogger(cfg.Detail, cfg.LogFile)
				if err != nil {
					return err
				}
				return http.RunServer(cfg, log)
			}
			return errors.New("missing config file")
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file (default is $config.yaml)")
}
