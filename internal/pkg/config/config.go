package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	NoteDir  string
	IndexDir string
}

func LoadConfig() {
	viper.SetConfigName("notington.yaml")
	viper.AddConfigPath("$HOME/.config/notington")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

func SetDefaults() {
	viper.SetDefault("NoteDir", "$HOME/notington")
	viper.SetDefault("IndexDir", "$HOME/.config/notington")
}
