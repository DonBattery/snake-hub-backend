package app

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/donbattery/snake-hub-backend/config"
	"github.com/donbattery/snake-hub-backend/logger"
	"github.com/donbattery/snake-hub-backend/server"
)

var (
	log     = logger.New("Server")
	conf    = config.New()
	command = &cobra.Command{}
)

// Run sets up and runs the server
func Run() {
	log.Info("Snake-Hub server initializing ...")
	setupCommand()
	if err := command.Execute(); err != nil {
		log.WithError(err).Fatal("Snake-hub server halted")
	}
}

func setupCommand() {
	command.RunE = setupAndRunServer
	command.PersistentFlags().StringVarP(&conf.ConfigPath, "config-file", "c", "", "Configuration file path")
	command.PersistentFlags().IntP("port", "p", 6767, "Snake-hub tcp port")
	command.PersistentFlags().StringP("snake_header", "H", "", "Auth Header Key")
	command.PersistentFlags().StringP("snake_secret", "S", "", "Auth Header Value")
}

func setupAndRunServer(cmd *cobra.Command, args []string) error {
	if err := conf.Init(cmd, args); err != nil {
		return errors.Wrap(err, "Failed to set up configurations")
	}
	log.WithField("PORT", viper.GetInt("port")).Info("HTTP Server listening")
	if err := server.Start(); err != nil {
		return errors.Wrap(err, "The HTTP server exited with an error")
	}
	return nil
}
