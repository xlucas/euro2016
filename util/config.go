package util

import "github.com/spf13/viper"

func LoadConfig() error {
	// Defaults
	viper.SetDefault("Token", "123")

	// Config path
	viper.SetConfigName(".euro2016")
	viper.AddConfigPath("$HOME")

	// Read config
	return viper.ReadInConfig()
}
