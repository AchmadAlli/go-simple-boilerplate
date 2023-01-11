package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type appENV struct {
	Name       string `mapstructure:"APP_NAME"`
	Port       int    `mapstructure:"APP_PORT"`
	Address    string `mapstructure:"APP_ADDRESS"`
	Identifier string `mapstructure:"APP_IDENTIFIER"`
	ENV        string `mapstructure:"APP_ENV"`
	Debug      bool   `mapstructure:"APP_DEBUG"`
}

var env = new(appENV)

func LoadEnv() {
	if _, err := os.Stat(".env.development"); err == nil {
		viper.AddConfigPath("./")
		viper.SetConfigName(".env.development")
		viper.SetConfigType("env")
		readConfigFile()
	} else {
		loadOsEnv()
	}

	if err := viper.Unmarshal(env); err != nil {
		panic(fmt.Sprintf("unable to parse env: %v", err))
	}
}

func GetEnv() *appENV {
	if env != nil {
		return env
	}

	LoadEnv()
	return env
}

func readConfigFile() {
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("unable to read env: %v", err))
	}
}

func loadOsEnv() {
	viper.BindEnv("APP_NAME")
	viper.BindEnv("APP_PORT")
	viper.BindEnv("APP_ADDRESS")
	viper.BindEnv("APP_IDENTIFIER")
	viper.BindEnv("APP_ENV")
	viper.BindEnv("APP_DEBUG")
}
