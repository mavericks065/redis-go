package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

// Config is loaded from the '.env' file if present and/or
// from environment variables.
// If both are set, environment variables take precedence.
const defaultConfigFile = ".env"

var config Config
var once sync.Once

type Config struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

func GetConfig() Config {
	once.Do(func() {
		// Load config from file if present
		if _, err := os.Stat(defaultConfigFile); err == nil {
			viper.SetConfigFile(defaultConfigFile)
			err := viper.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("fatal error config file: %s", err))
			}
		}

		// Load config from env variables
		viper.BindEnv("HOST")
		viper.BindEnv("PORT")
		viper.AutomaticEnv()

		err := viper.Unmarshal(&config)
		if err != nil {
			panic(fmt.Errorf("fatal error config parsing: %s", err))
		}
	})

	return config
}
