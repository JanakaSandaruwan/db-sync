package db

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// Unmarshal the config into a struct
type Config struct {
	SourceDB DBConfig `mapstructure:"source_db"`
	TargetDB DBConfig `mapstructure:"target_db"`
}

var config Config

func loadConfig(cfgFile string) (Config, error) {
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()

	if err != nil {
		return Config{}, fmt.Errorf("fatal error while reading config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return config, nil
}
