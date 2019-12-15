package config

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/donbattery/snake-hub-backend/logger"
	"github.com/donbattery/snake-hub-backend/model"
)

// Config is the main configuration object
type Config struct {
	// ConfigPath is the path to the config file
	ConfigPath string
	once       sync.Once
	log        *logger.Logger
	gameConfig *model.GameConfig
}

// New returns a new Snake-hub configuration object
func New() *Config {
	return &Config{
		log: logger.New("Config"),
	}
}

// Init sets up Viper according to defaults, configfile, envs and Cobra flags
func (conf *Config) Init(cmd *cobra.Command, args []string) error {
	// Bind Cobra flags to Viper keys
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return errors.Wrap(err, "Failed to bind Cobra flags to Viper keys")
	}
	// Search the environment for variables prefixed with PAGER_
	viper.SetEnvPrefix("SNAKE_HUB")
	// Read in environment variables that match
	viper.AutomaticEnv()
	// If a ConfigFile is set read it in
	if conf.ConfigPath != "" {
		// Use config file from the flag
		viper.SetConfigFile(conf.ConfigPath)
		if err := viper.ReadInConfig(); err != nil {
			return errors.Wrapf(err, "Failed to load Viper configs from file %s", conf.ConfigPath)
		}
	}
	// Unmarshal Game configs
	if err := viper.UnmarshalKey("game_config", &conf.gameConfig); err != nil {
		return errors.Wrap(err, "Failed to unmarshal games configs")
	}

	return nil
}
